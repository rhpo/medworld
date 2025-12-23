<?php

namespace App\Http\Controllers\Api;

use App\Models\Message;
use Illuminate\Http\Request;

class MessageController
{
    public function index()
    {
        $messages = Message::with(['sender', 'receiver', 'cabinet'])->paginate(15);
        return response()->json($messages, 200);
    }

    public function show($id)
    {
        $message = Message::with(['sender', 'receiver', 'cabinet'])->find($id);

        if (!$message) {
            return response()->json(['message' => 'Message not found'], 404);
        }

        return response()->json($message, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'sender_id' => 'required|exists:users,id',
            'receiver_id' => 'required|exists:users,id',
            'cabinet_id' => 'required|exists:cabinets,id',
            'content' => 'required|array',
            'status' => 'required|in:unseen,seen',
            'attachments' => 'nullable|array',
        ]);

        $message = Message::create($validated);

        return response()->json([
            'message' => 'Message created successfully',
            'data' => $message->load(['sender', 'receiver', 'cabinet']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $message = Message::find($id);

        if (!$message) {
            return response()->json(['message' => 'Message not found'], 404);
        }

        $validated = $request->validate([
            'content' => 'sometimes|array',
            'status' => 'sometimes|in:unseen,seen',
            'attachments' => 'nullable|array',
        ]);

        $message->update($validated);

        return response()->json([
            'message' => 'Message updated successfully',
            'data' => $message->load(['sender', 'receiver', 'cabinet']),
        ], 200);
    }

    public function destroy($id)
    {
        $message = Message::find($id);

        if (!$message) {
            return response()->json(['message' => 'Message not found'], 404);
        }

        $message->delete();

        return response()->json(['message' => 'Message deleted successfully'], 200);
    }

    public function getConversation($userId1, $userId2)
    {
        $messages = Message::where(function ($query) use ($userId1, $userId2) {
            $query->where('sender_id', $userId1)->where('receiver_id', $userId2);
        })->orWhere(function ($query) use ($userId1, $userId2) {
            $query->where('sender_id', $userId2)->where('receiver_id', $userId1);
        })->with(['sender', 'receiver', 'cabinet'])->orderBy('created_at', 'asc')->paginate(50);

        return response()->json($messages, 200);
    }

    public function getUserMessages($userId)
    {
        $messages = Message::where(function ($query) use ($userId) {
            $query->where('sender_id', $userId)->orWhere('receiver_id', $userId);
        })->with(['sender', 'receiver', 'cabinet'])->orderBy('created_at', 'desc')->paginate(15);

        return response()->json($messages, 200);
    }

    public function markAsSeen(Request $request, $id)
    {
        $message = Message::find($id);

        if (!$message) {
            return response()->json(['message' => 'Message not found'], 404);
        }

        $message->update(['status' => 'seen']);

        return response()->json([
            'message' => 'Message marked as seen',
            'data' => $message->load(['sender', 'receiver', 'cabinet']),
        ], 200);
    }
}
