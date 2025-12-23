<?php

namespace App\Http\Controllers\Api;

use App\Models\Assistant;
use Illuminate\Http\Request;
use App\Models\User;
use Illuminate\Support\Facades\Hash;

class AssistantController
{
    public function index()
    {
        $assistants = Assistant::with(['user', 'cabinet', 'doctors'])->paginate(15);
        return response()->json($assistants, 200);
    }

    public function show($id)
    {
        $assistant = Assistant::with(['user', 'cabinet', 'doctors'])->find($id);

        if (!$assistant) {
            return response()->json(['message' => 'Assistant not found'], 404);
        }

        return response()->json($assistant, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'user_id' => 'required|exists:users,id',
            'cabinet_id' => 'required|exists:cabinets,id',
        ]);

        $assistant = Assistant::create($validated);

        return response()->json([
            'message' => 'Assistant created successfully',
            'assistant' => $assistant->load(['user', 'cabinet', 'doctors']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $assistant = Assistant::find($id);

        if (!$assistant) {
            return response()->json(['message' => 'Assistant not found'], 404);
        }

        $validated = $request->validate([
            'cabinet_id' => 'sometimes|exists:cabinets,id',
        ]);

        $assistant->update($validated);

        return response()->json([
            'message' => 'Assistant updated successfully',
            'assistant' => $assistant->load(['user', 'cabinet', 'doctors']),
        ], 200);
    }

    public function destroy($id)
    {
        $assistant = Assistant::find($id);

        if (!$assistant) {
            return response()->json(['message' => 'Assistant not found'], 404);
        }

        $assistant->delete();

        return response()->json(['message' => 'Assistant deleted successfully'], 200);
    }

    public function getDoctors($id)
    {
        $assistant = Assistant::find($id);

        if (!$assistant) {
            return response()->json(['message' => 'Assistant not found'], 404);
        }

        $doctors = $assistant->doctors()->with(['user', 'cabinet'])->paginate(15);

        return response()->json($doctors, 200);
    }

    public function attachDoctor(Request $request, $id)
    {
        $assistant = Assistant::find($id);

        if (!$assistant) {
            return response()->json(['message' => 'Assistant not found'], 404);
        }

        $validated = $request->validate([
            'doctor_id' => 'required|exists:doctors,id',
        ]);

        $assistant->doctors()->attach($validated['doctor_id']);

        return response()->json([
            'message' => 'Doctor attached to assistant',
            'assistant' => $assistant->load(['doctors']),
        ], 200);
    }

    public function detachDoctor(Request $request, $id)
    {
        $assistant = Assistant::find($id);

        if (!$assistant) {
            return response()->json(['message' => 'Assistant not found'], 404);
        }

        $validated = $request->validate([
            'doctor_id' => 'required|exists:doctors,id',
        ]);

        $assistant->doctors()->detach($validated['doctor_id']);

        return response()->json([
            'message' => 'Doctor detached from assistant',
            'assistant' => $assistant->load(['doctors']),
        ], 200);
    }
}
