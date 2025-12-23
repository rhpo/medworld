<?php

namespace App\Http\Controllers\Api;

use App\Models\Prescription;
use Illuminate\Http\Request;

class PrescriptionController
{
    public function index()
    {
        $prescriptions = Prescription::with(['consultation', 'patient', 'doctor'])->paginate(15);
        return response()->json($prescriptions, 200);
    }

    public function show($id)
    {
        $prescription = Prescription::with(['consultation', 'patient', 'doctor'])->find($id);

        if (!$prescription) {
            return response()->json(['message' => 'Prescription not found'], 404);
        }

        return response()->json($prescription, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'consultation_id' => 'required|exists:consultations,id',
            'patient_id' => 'required|exists:patients,id',
            'doctor_id' => 'required|exists:doctors,id',
            'prescription_date' => 'required|date',
            'status' => 'required|in:ACTIVE,COMPLETED,CANCELLED',
            'medications' => 'nullable|array',
            'general_instructions' => 'nullable|string',
            'valid_until' => 'nullable|date',
            'refills_allowed' => 'nullable|integer|min:0',
            'refills_used' => 'nullable|integer|min:0',
        ]);

        $prescription = Prescription::create($validated);

        return response()->json([
            'message' => 'Prescription created successfully',
            'prescription' => $prescription->load(['consultation', 'patient', 'doctor']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $prescription = Prescription::find($id);

        if (!$prescription) {
            return response()->json(['message' => 'Prescription not found'], 404);
        }

        $validated = $request->validate([
            'status' => 'sometimes|in:ACTIVE,COMPLETED,CANCELLED',
            'medications' => 'nullable|array',
            'general_instructions' => 'nullable|string',
            'valid_until' => 'nullable|date',
            'refills_allowed' => 'nullable|integer|min:0',
            'refills_used' => 'nullable|integer|min:0',
        ]);

        $prescription->update($validated);

        return response()->json([
            'message' => 'Prescription updated successfully',
            'prescription' => $prescription->load(['consultation', 'patient', 'doctor']),
        ], 200);
    }

    public function destroy($id)
    {
        $prescription = Prescription::find($id);

        if (!$prescription) {
            return response()->json(['message' => 'Prescription not found'], 404);
        }

        $prescription->delete();

        return response()->json(['message' => 'Prescription deleted successfully'], 200);
    }

    public function getByPatient($patientId)
    {
        $prescriptions = Prescription::where('patient_id', $patientId)
            ->with(['doctor', 'consultation'])
            ->paginate(15);

        return response()->json($prescriptions, 200);
    }

    public function getByDoctor($doctorId)
    {
        $prescriptions = Prescription::where('doctor_id', $doctorId)
            ->with(['patient', 'consultation'])
            ->paginate(15);

        return response()->json($prescriptions, 200);
    }
}
