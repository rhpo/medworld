<?php

namespace App\Http\Controllers\Api;

use App\Models\Doctor;
use App\Models\Patient;
use App\Models\Appointment;
use App\Models\Assistant;
use App\Models\Cabinet;
use App\Models\Consultation;
use App\Models\User;
use Illuminate\Http\Request;

class AllController
{
    public function listAllDoctors()
    {
        $doctors = Doctor::with(['user', 'cabinet', 'assistants'])->paginate(50);
        return response()->json($doctors, 200);
    }

    public function listAllPatients()
    {
        $patients = Patient::with(['user', 'appointments', 'consultations'])->paginate(50);
        return response()->json($patients, 200);
    }

    public function listAllAppointments()
    {
        $appointments = Appointment::with([
            'patient.user',
            'doctor.user',
            'cabinet',
            'consultation',
            'createdByAssistant',
            'payment'
        ])->paginate(50);
        return response()->json($appointments, 200);
    }

    public function listAllAssistants()
    {
        $assistants = Assistant::with(['user', 'cabinet', 'doctors'])->paginate(50);
        return response()->json($assistants, 200);
    }

    public function listAllCabinets()
    {
        $cabinets = Cabinet::with([
            'admin',
            'doctors.user',
            'assistants.user',
            'appointments',
            'ratings'
        ])->paginate(50);
        return response()->json($cabinets, 200);
    }

    public function listAllConsultations()
    {
        $consultations = Consultation::with([
            'doctor.user',
            'patient.user',
            'appointment'
        ])->paginate(50);
        return response()->json($consultations, 200);
    }

    public function listAllUsers()
    {
        $users = User::with(['doctor', 'patient'])->paginate(50);
        return response()->json($users, 200);
    }
}
