<?php

namespace App\Http\Controllers\Api;

use App\Models\Consultation;
use Illuminate\Http\Request;

class ConsultationController
{
    public function index()
    {
        $consultations = Consultation::with(['doctor', 'patient', 'appointment'])->paginate(15);
        return response()->json($consultations, 200);
    }

    public function show($id)
    {
        $consultation = Consultation::with(['doctor', 'patient', 'appointment'])->find($id);

        if (!$consultation) {
            return response()->json(['message' => 'Consultation not found'], 404);
        }

        return response()->json($consultation, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'appointment_id' => 'required|exists:appointments,id|unique:consultations,appointment_id',
            'doctor_id' => 'required|exists:doctors,id',
            'patient_id' => 'required|exists:patients,id',
            'notes' => 'nullable|string',
            'prescriptions' => 'nullable|array',
            'attachments' => 'nullable|array',
        ]);

        $consultation = Consultation::create($validated);

        return response()->json([
            'message' => 'Consultation created successfully',
            'consultation' => $consultation->load(['doctor', 'patient', 'appointment']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $consultation = Consultation::find($id);

        if (!$consultation) {
            return response()->json(['message' => 'Consultation not found'], 404);
        }

        $validated = $request->validate([
            'notes' => 'nullable|string',
            'prescriptions' => 'nullable|array',
            'attachments' => 'nullable|array',
        ]);

        $consultation->update($validated);

        return response()->json([
            'message' => 'Consultation updated successfully',
            'consultation' => $consultation->load(['doctor', 'patient', 'appointment']),
        ], 200);
    }

    public function destroy($id)
    {
        $consultation = Consultation::find($id);

        if (!$consultation) {
            return response()->json(['message' => 'Consultation not found'], 404);
        }

        $consultation->delete();

        return response()->json(['message' => 'Consultation deleted successfully'], 200);
    }

    public function getByPatient($patientId)
    {
        $consultations = Consultation::where('patient_id', $patientId)
            ->with(['doctor', 'appointment'])
            ->paginate(15);

        return response()->json($consultations, 200);
    }

    public function getByDoctor($doctorId)
    {
        $consultations = Consultation::where('doctor_id', $doctorId)
            ->with(['patient', 'appointment'])
            ->paginate(15);

        return response()->json($consultations, 200);
    }
}
