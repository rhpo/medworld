<?php

namespace App\Http\Controllers\Api;

use App\Models\Doctor;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;
use App\Models\User;

class DoctorController
{
    public function index()
    {
        $doctors = Doctor::with(['user', 'cabinet', 'assistants', 'appointments', 'consultations', 'payments'])->paginate(15);
        return response()->json($doctors, 200);
    }

    public function show($id)
    {
        $doctor = Doctor::with(['user', 'cabinet', 'assistants', 'appointments', 'consultations', 'payments'])->find($id);

        if (!$doctor) {
            return response()->json(['message' => 'Doctor not found'], 404);
        }

        return response()->json($doctor, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'user_id' => 'required|exists:users,id',
            'cabinet_id' => 'required|exists:cabinets,id',
            'speciality' => 'required|string|max:255',
            'career_start' => 'required|date',
            'consultation_price' => 'required|numeric|min:0',
            'consultation_duration' => 'required|integer|min:5',
        ]);

        $doctor = Doctor::create($validated);

        return response()->json([
            'message' => 'Doctor created successfully',
            'doctor' => $doctor->load(['user', 'cabinet', 'assistants']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $doctor = Doctor::find($id);

        if (!$doctor) {
            return response()->json(['message' => 'Doctor not found'], 404);
        }

        $validated = $request->validate([
            'cabinet_id' => 'sometimes|exists:cabinets,id',
            'speciality' => 'sometimes|string|max:255',
            'career_start' => 'sometimes|date',
            'consultation_price' => 'sometimes|numeric|min:0',
            'consultation_duration' => 'sometimes|integer|min:5',
        ]);

        $doctor->update($validated);

        return response()->json([
            'message' => 'Doctor updated successfully',
            'doctor' => $doctor->load(['user', 'cabinet', 'assistants']),
        ], 200);
    }

    public function destroy($id)
    {
        $doctor = Doctor::find($id);

        if (!$doctor) {
            return response()->json(['message' => 'Doctor not found'], 404);
        }

        $doctor->delete();

        return response()->json(['message' => 'Doctor deleted successfully'], 200);
    }

    public function getAppointments($id)
    {
        $doctor = Doctor::find($id);

        if (!$doctor) {
            return response()->json(['message' => 'Doctor not found'], 404);
        }

        $appointments = $doctor->appointments()->with(['patient', 'cabinet', 'consultation'])->paginate(15);

        return response()->json($appointments, 200);
    }

    public function getConsultations($id)
    {
        $doctor = Doctor::find($id);

        if (!$doctor) {
            return response()->json(['message' => 'Doctor not found'], 404);
        }

        $consultations = $doctor->consultations()->with(['patient', 'appointment'])->paginate(15);

        return response()->json($consultations, 200);
    }

    public function getAssistants($id)
    {
        $doctor = Doctor::find($id);

        if (!$doctor) {
            return response()->json(['message' => 'Doctor not found'], 404);
        }

        $assistants = $doctor->assistants()->with(['user', 'cabinet'])->paginate(15);

        return response()->json($assistants, 200);
    }

    public function getPatients($id)
    {
        $doctor = Doctor::find($id);

        if (!$doctor) {
            return response()->json(['message' => 'Doctor not found'], 404);
        }

        // Get all unique patients who have appointments with this doctor
        $patients = \App\Models\Patient::whereHas('appointments', function ($query) use ($id) {
            $query->where('doctor_id', $id);
        })->with(['user', 'appointments', 'consultations'])->paginate(15);

        return response()->json($patients, 200);
    }

    /**
     * Search/Filter doctors by criteria
     * Supports filtering by: speciality, cabinet_id, price_min, price_max, availability
     *
     * Query parameters:
     * - speciality: string (e.g., "Cardiology")
     * - cabinet_id: integer
     * - price_max: integer (max consultation price)
     * - price_min: integer (min consultation price)
     * - available: boolean (doctors with available appointments)
     */
    public function search(Request $request)
    {
        $query = Doctor::query()->with(['user', 'cabinet', 'assistants', 'appointments', 'consultations']);

        // Filter by speciality
        if ($request->has('speciality') && $request->speciality) {
            $query->where('speciality', $request->speciality);
        }

        // Filter by cabinet
        if ($request->has('cabinet_id') && $request->cabinet_id) {
            $query->where('cabinet_id', $request->cabinet_id);
        }

        // Filter by max price
        if ($request->has('price_max') && $request->price_max) {
            $query->where('consultation_price', '<=', $request->price_max);
        }

        // Filter by min price
        if ($request->has('price_min') && $request->price_min) {
            $query->where('consultation_price', '>=', $request->price_min);
        }

        // Filter by availability (has open appointment slots or no appointments yet)
        if ($request->has('available') && $request->available) {
            // Doctors with fewer than 20 appointments or none
            $query->withCount('appointments')
                ->having('appointments_count', '<', 20);
        }

        $doctors = $query->paginate(15);

        return response()->json($doctors, 200);
    }
}
