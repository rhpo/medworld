<?php

namespace App\Http\Controllers\Api;

use App\Models\Cabinet;
use Illuminate\Http\Request;

class CabinetController
{
    public function index()
    {
        $cabinets = Cabinet::with(['admin', 'doctors', 'assistants', 'appointments', 'messages', 'ratings', 'payments'])->paginate(15);
        return response()->json($cabinets, 200);
    }

    public function show($id)
    {
        $cabinet = Cabinet::with(['admin', 'doctors', 'assistants', 'appointments', 'messages', 'ratings', 'payments'])->find($id);

        if (!$cabinet) {
            return response()->json(['message' => 'Cabinet not found'], 404);
        }

        return response()->json($cabinet, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'user_id' => 'required|exists:users,id',
            'name' => 'required|string|max:255',
            'phone' => 'required|string|max:20',
            'opening_hours' => 'nullable|array',
            'location_lat' => 'nullable|numeric',
            'location_lng' => 'nullable|numeric',
            'access_handicap' => 'nullable|boolean',
            'has_parking' => 'nullable|boolean',
            'has_wifi' => 'nullable|boolean',
            'accepts_urgent' => 'nullable|boolean',
            'accepts_insurance' => 'nullable|boolean',
        ]);

        $cabinet = Cabinet::create($validated);

        return response()->json([
            'message' => 'Cabinet created successfully',
            'cabinet' => $cabinet->load(['admin', 'doctors', 'assistants']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $cabinet = Cabinet::find($id);

        if (!$cabinet) {
            return response()->json(['message' => 'Cabinet not found'], 404);
        }

        $validated = $request->validate([
            'name' => 'sometimes|string|max:255',
            'phone' => 'sometimes|string|max:20',
            'opening_hours' => 'nullable|array',
            'location_lat' => 'nullable|numeric',
            'location_lng' => 'nullable|numeric',
            'access_handicap' => 'nullable|boolean',
            'has_parking' => 'nullable|boolean',
            'has_wifi' => 'nullable|boolean',
            'accepts_urgent' => 'nullable|boolean',
            'accepts_insurance' => 'nullable|boolean',
        ]);

        $cabinet->update($validated);

        return response()->json([
            'message' => 'Cabinet updated successfully',
            'cabinet' => $cabinet->load(['admin', 'doctors', 'assistants']),
        ], 200);
    }

    public function destroy($id)
    {
        $cabinet = Cabinet::find($id);

        if (!$cabinet) {
            return response()->json(['message' => 'Cabinet not found'], 404);
        }

        $cabinet->delete();

        return response()->json(['message' => 'Cabinet deleted successfully'], 200);
    }

    public function getDoctors($id)
    {
        $cabinet = Cabinet::find($id);

        if (!$cabinet) {
            return response()->json(['message' => 'Cabinet not found'], 404);
        }

        $doctors = $cabinet->doctors()->with(['user', 'assistants', 'appointments'])->paginate(15);

        return response()->json($doctors, 200);
    }

    public function getAssistants($id)
    {
        $cabinet = Cabinet::find($id);

        if (!$cabinet) {
            return response()->json(['message' => 'Cabinet not found'], 404);
        }

        $assistants = $cabinet->assistants()->with(['user', 'doctors'])->paginate(15);

        return response()->json($assistants, 200);
    }

    public function getAppointments($id)
    {
        $cabinet = Cabinet::find($id);

        if (!$cabinet) {
            return response()->json(['message' => 'Cabinet not found'], 404);
        }

        $appointments = $cabinet->appointments()->with(['patient', 'doctor', 'consultation'])->paginate(15);

        return response()->json($appointments, 200);
    }

    public function getRatings($id)
    {
        $cabinet = Cabinet::find($id);

        if (!$cabinet) {
            return response()->json(['message' => 'Cabinet not found'], 404);
        }

        $ratings = $cabinet->ratings()->with(['patient'])->paginate(15);

        return response()->json($ratings, 200);
    }
}
