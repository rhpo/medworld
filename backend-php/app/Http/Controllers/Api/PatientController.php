<?php

namespace App\Http\Controllers\Api;

use App\Models\Patient;
use Illuminate\Http\Request;
use App\Models\User;
use Illuminate\Support\Facades\Hash;

class PatientController
{
    public function index()
    {
        $patients = Patient::with(['user', 'appointments', 'consultations', 'prescriptions', 'ratings', 'payments'])->paginate(15);
        return response()->json($patients, 200);
    }

    public function show($id)
    {
        $patient = Patient::with(['user', 'appointments', 'consultations', 'prescriptions', 'ratings', 'payments'])->find($id);

        if (!$patient) {
            return response()->json(['message' => 'Patient not found'], 404);
        }

        return response()->json($patient, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'user_id' => 'required|exists:users,id',
            'blood_type' => 'nullable|string|max:5',
            'medical_history' => 'nullable|array',
            'allergies' => 'nullable|array',
            'emergency_contact' => 'nullable|string|max:255',
            'weight' => 'nullable|numeric|min:0',
        ]);

        $patient = Patient::create($validated);

        return response()->json([
            'message' => 'Patient created successfully',
            'patient' => $patient->load(['user', 'appointments', 'consultations']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $patient = Patient::find($id);

        if (!$patient) {
            return response()->json(['message' => 'Patient not found'], 404);
        }

        $validated = $request->validate([
            'blood_type' => 'nullable|string|max:5',
            'medical_history' => 'nullable|array',
            'allergies' => 'nullable|array',
            'emergency_contact' => 'nullable|string|max:255',
            'weight' => 'nullable|numeric|min:0',
        ]);

        $patient->update($validated);

        return response()->json([
            'message' => 'Patient updated successfully',
            'patient' => $patient->load(['user', 'appointments', 'consultations']),
        ], 200);
    }

    public function destroy($id)
    {
        $patient = Patient::find($id);

        if (!$patient) {
            return response()->json(['message' => 'Patient not found'], 404);
        }

        $patient->delete();

        return response()->json(['message' => 'Patient deleted successfully'], 200);
    }

    public function getAppointments($id)
    {
        $patient = Patient::find($id);

        if (!$patient) {
            return response()->json(['message' => 'Patient not found'], 404);
        }

        $appointments = $patient->appointments()->with(['doctor', 'cabinet', 'consultation'])->paginate(15);

        return response()->json($appointments, 200);
    }

    public function getConsultations($id)
    {
        $patient = Patient::find($id);

        if (!$patient) {
            return response()->json(['message' => 'Patient not found'], 404);
        }

        $consultations = $patient->consultations()->with(['doctor', 'appointment'])->paginate(15);

        return response()->json($consultations, 200);
    }

    public function getPrescriptions($id)
    {
        $patient = Patient::find($id);

        if (!$patient) {
            return response()->json(['message' => 'Patient not found'], 404);
        }

        $prescriptions = $patient->prescriptions()->with(['doctor', 'consultation'])->paginate(15);

        return response()->json($prescriptions, 200);
    }
}
