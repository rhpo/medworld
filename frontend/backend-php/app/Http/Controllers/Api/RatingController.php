<?php

namespace App\Http\Controllers\Api;

use App\Models\Rating;
use Illuminate\Http\Request;

class RatingController
{
    public function index()
    {
        $ratings = Rating::with(['patient', 'cabinet'])->paginate(15);
        return response()->json($ratings, 200);
    }

    public function show($id)
    {
        $rating = Rating::with(['patient', 'cabinet'])->find($id);

        if (!$rating) {
            return response()->json(['message' => 'Rating not found'], 404);
        }

        return response()->json($rating, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'patient_id' => 'required|exists:patients,id',
            'cabinet_id' => 'required|exists:cabinets,id',
            'date' => 'required|date',
            'equippement' => 'nullable|array',
            'user_experience' => 'nullable|array',
            'review' => 'nullable|string',
        ]);

        $rating = Rating::create($validated);

        return response()->json([
            'message' => 'Rating created successfully',
            'rating' => $rating->load(['patient', 'cabinet']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $rating = Rating::find($id);

        if (!$rating) {
            return response()->json(['message' => 'Rating not found'], 404);
        }

        $validated = $request->validate([
            'equippement' => 'nullable|array',
            'user_experience' => 'nullable|array',
            'review' => 'nullable|string',
        ]);

        $rating->update($validated);

        return response()->json([
            'message' => 'Rating updated successfully',
            'rating' => $rating->load(['patient', 'cabinet']),
        ], 200);
    }

    public function destroy($id)
    {
        $rating = Rating::find($id);

        if (!$rating) {
            return response()->json(['message' => 'Rating not found'], 404);
        }

        $rating->delete();

        return response()->json(['message' => 'Rating deleted successfully'], 200);
    }

    public function getByCabinet($cabinetId)
    {
        $ratings = Rating::where('cabinet_id', $cabinetId)
            ->with(['patient'])
            ->paginate(15);

        return response()->json($ratings, 200);
    }

    public function getByPatient($patientId)
    {
        $ratings = Rating::where('patient_id', $patientId)
            ->with(['cabinet'])
            ->paginate(15);

        return response()->json($ratings, 200);
    }
}
