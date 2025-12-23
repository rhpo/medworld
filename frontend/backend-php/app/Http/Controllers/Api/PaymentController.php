<?php

namespace App\Http\Controllers\Api;

use App\Models\Payment;
use Illuminate\Http\Request;

class PaymentController
{
    public function index()
    {
        $payments = Payment::with(['patient', 'doctor', 'cabinet', 'appointment'])->paginate(15);
        return response()->json($payments, 200);
    }

    public function show($id)
    {
        $payment = Payment::with(['patient', 'doctor', 'cabinet', 'appointment'])->find($id);

        if (!$payment) {
            return response()->json(['message' => 'Payment not found'], 404);
        }

        return response()->json($payment, 200);
    }

    public function store(Request $request)
    {
        $validated = $request->validate([
            'patient_id' => 'required|exists:patients,id',
            'doctor_id' => 'required|exists:doctors,id',
            'cabinet_id' => 'required|exists:cabinets,id',
            'appointment_id' => 'nullable|exists:appointments,id',
            'amount' => 'required|numeric|min:0',
            'status' => 'required|in:pending,completed,failed,refunded',
            'payment_method' => 'required|string|max:50',
            'transaction_date' => 'required|date',
            'notes' => 'nullable|string',
        ]);

        $payment = Payment::create($validated);

        return response()->json([
            'message' => 'Payment created successfully',
            'payment' => $payment->load(['patient', 'doctor', 'cabinet', 'appointment']),
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $payment = Payment::find($id);

        if (!$payment) {
            return response()->json(['message' => 'Payment not found'], 404);
        }

        $validated = $request->validate([
            'amount' => 'sometimes|numeric|min:0',
            'status' => 'sometimes|in:pending,completed,failed,refunded',
            'payment_method' => 'sometimes|string|max:50',
            'transaction_date' => 'sometimes|date',
            'notes' => 'nullable|string',
        ]);

        $payment->update($validated);

        return response()->json([
            'message' => 'Payment updated successfully',
            'payment' => $payment->load(['patient', 'doctor', 'cabinet', 'appointment']),
        ], 200);
    }

    public function destroy($id)
    {
        $payment = Payment::find($id);

        if (!$payment) {
            return response()->json(['message' => 'Payment not found'], 404);
        }

        $payment->delete();

        return response()->json(['message' => 'Payment deleted successfully'], 200);
    }

    public function getByPatient($patientId)
    {
        $payments = Payment::where('patient_id', $patientId)
            ->with(['doctor', 'cabinet', 'appointment'])
            ->paginate(15);

        return response()->json($payments, 200);
    }

    public function getByDoctor($doctorId)
    {
        $payments = Payment::where('doctor_id', $doctorId)
            ->with(['patient', 'cabinet', 'appointment'])
            ->paginate(15);

        return response()->json($payments, 200);
    }

    public function getByCabinet($cabinetId)
    {
        $payments = Payment::where('cabinet_id', $cabinetId)
            ->with(['patient', 'doctor', 'appointment'])
            ->paginate(15);

        return response()->json($payments, 200);
    }

    public function getByStatus($status)
    {
        $payments = Payment::where('status', $status)
            ->with(['patient', 'doctor', 'cabinet', 'appointment'])
            ->paginate(15);

        return response()->json($payments, 200);
    }
}
