<?php

namespace App\Http\Controllers\Api;

use App\Models\Appointment;
use Illuminate\Http\Request;

class AppointmentController
{
    public function index()
    {
        $appointments = Appointment::with(['patient', 'doctor', 'cabinet', 'consultation', 'createdByAssistant', 'payment'])->paginate(15);
        return response()->json($appointments, 200);
    }

    public function show($id)
    {
        $appointment = Appointment::with(['patient', 'doctor', 'cabinet', 'consultation', 'createdByAssistant', 'payment'])->find($id);

        if (!$appointment) {
            return response()->json(['message' => 'Appointment not found'], 404);
        }

        return response()->json($appointment, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'patient_id' => 'required|exists:patients,id',
            'doctor_id' => 'required|exists:doctors,id',
            'cabinet_id' => 'required|exists:cabinets,id',
            'date' => 'required|date_format:Y-m-d H:i:s',
            'status' => 'required|in:SCHEDULED,CONFIRMED,IN_PROGRESS,COMPLETED,CANCELLED,NO_SHOW',
            'created_by_assistant_id' => 'nullable|exists:assistants,id',
        ]);

        $appointment = Appointment::create($validated);

        return response()->json([
            'message' => 'Appointment created successfully',
            'appointment' => $appointment->load(['patient', 'doctor', 'cabinet', 'createdByAssistant']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $appointment = Appointment::find($id);

        if (!$appointment) {
            return response()->json(['message' => 'Appointment not found'], 404);
        }

        $validated = $request->validate([
            'date' => 'sometimes|date_format:Y-m-d H:i:s',
            'status' => 'sometimes|in:SCHEDULED,CONFIRMED,IN_PROGRESS,COMPLETED,CANCELLED,NO_SHOW',
            'created_by_assistant_id' => 'nullable|exists:assistants,id',
        ]);

        $appointment->update($validated);

        return response()->json([
            'message' => 'Appointment updated successfully',
            'appointment' => $appointment->load(['patient', 'doctor', 'cabinet', 'createdByAssistant']),
        ], 200);
    }

    public function destroy($id)
    {
        $appointment = Appointment::find($id);

        if (!$appointment) {
            return response()->json(['message' => 'Appointment not found'], 404);
        }

        $appointment->delete();

        return response()->json(['message' => 'Appointment deleted successfully'], 200);
    }

    public function getByPatient($patientId)
    {
        $appointments = Appointment::where('patient_id', $patientId)
            ->with(['doctor', 'cabinet', 'consultation', 'createdByAssistant', 'payment'])
            ->paginate(15);

        return response()->json($appointments, 200);
    }

    public function getByDoctor($doctorId)
    {
        $appointments = Appointment::where('doctor_id', $doctorId)
            ->with(['patient', 'cabinet', 'consultation', 'createdByAssistant', 'payment'])
            ->paginate(15);

        return response()->json($appointments, 200);
    }

    public function getByCabinet($cabinetId)
    {
        $appointments = Appointment::where('cabinet_id', $cabinetId)
            ->with(['patient', 'doctor', 'consultation', 'createdByAssistant', 'payment'])
            ->paginate(15);

        return response()->json($appointments, 200);
    }
}
