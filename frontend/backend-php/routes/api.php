<?php

use Illuminate\Support\Facades\Route;
use Illuminate\Http\Request;
use App\Http\Controllers\Api\UserController;
use App\Http\Controllers\Api\AuthController;
use App\Http\Controllers\Api\RatingController;
use App\Http\Controllers\Api\DoctorController;
use App\Http\Controllers\Api\PatientController;
use App\Http\Controllers\Api\MessageController;
use App\Http\Controllers\Api\CabinetController;
use App\Http\Controllers\Api\PaymentController;
use App\Http\Controllers\Api\AssistantController;
use App\Http\Controllers\Api\AppointmentController;
use App\Http\Controllers\Api\ConsultationController;
use App\Http\Controllers\Api\PrescriptionController;
use App\Http\Controllers\Api\AllController;

Route::options('{any}', function (Request $request) {
    return response()->noContent();
})->where('any', '.*');

Route::prefix('v1')->group(function () {
    // PUBLIC Routes
    Route::get('/auth/login', function () {
        return "Please use POST!";
    });

    Route::post('auth/login', [AuthController::class, 'login']);
    Route::post('auth/register', [AuthController::class, 'register']);

    // Public doctor search endpoint - allows anyone to search for doctors
    Route::get('doctors/search/filter', [DoctorController::class, 'search']);

    // Public "All" endpoints for frontend - allow browsing doctors, cabinets before login
    Route::prefix('all')->group(function () {
        Route::get('doctors', [AllController::class, 'listAllDoctors']);
        Route::get('cabinets', [AllController::class, 'listAllCabinets']);
    });

    Route::middleware(['auth.cookie', 'auth:sanctum'])->group(function () {
        Route::post('/auth/logout', [AuthController::class, 'logout']);
        Route::get('/auth/me', [AuthController::class, 'me']);

        // User CRUD routes !! admin only
        Route::middleware('role:admin')->group(function () {
            Route::apiResource('users', UserController::class);
        });

        // Doctor routes Ddoctors and Admins
        Route::middleware('role:doctor,admin')->group(function () {
            Route::get('doctors/{id}/appointments', [DoctorController::class, 'getAppointments']);
            Route::get('doctors/{id}/consultations', [DoctorController::class, 'getConsultations']);
            Route::get('doctors/{id}/assistants', [DoctorController::class, 'getAssistants']);
            Route::get('doctors/{id}/patients', [DoctorController::class, 'getPatients']);
        });

        // Doctor CRUD admin only
        Route::middleware('role:admin')->group(function () {
            Route::apiResource('doctors', DoctorController::class);
        });

        // Patient routes Patients , Doctors, and Admins
        Route::middleware('role:patient,doctor,admin')->group(function () {
            Route::get('patients/{id}/appointments', [PatientController::class, 'getAppointments']);
            Route::get('patients/{id}/consultations', [PatientController::class, 'getConsultations']);
            Route::get('patients/{id}/prescriptions', [PatientController::class, 'getPrescriptions']);
        });

        // Patient CRUD - admin only
        Route::middleware('role:doctor,admin')->group(function () {
            Route::apiResource('patients', PatientController::class);
        });

        // Cabinet CRUD and routes - admin and Cabinet managers
        Route::middleware('role:superadmin,admin')->group(function () {
            Route::apiResource('cabinets', CabinetController::class);
            Route::get('cabinets/{id}/doctors', [CabinetController::class, 'getDoctors']);
            Route::get('cabinets/{id}/assistants', [CabinetController::class, 'getAssistants']);
            Route::get('cabinets/{id}/appointments', [CabinetController::class, 'getAppointments']);
            Route::get('cabinets/{id}/ratings', [CabinetController::class, 'getRatings']);
        });

        // Assistant CRUD  Admin and Doctors
        Route::middleware('role:admin,doctor')->group(function () {
            Route::apiResource('assistants', AssistantController::class);
            Route::get('assistants/{id}/doctors', [AssistantController::class, 'getDoctors']);
            Route::post('assistants/{id}/doctors/attach', [AssistantController::class, 'attachDoctor']);
            Route::post('assistants/{id}/doctors/detach', [AssistantController::class, 'detac   hDoctor']);
        });

        // Appointment routes - Available to all authenticated users
        Route::apiResource('appointments', AppointmentController::class);
        Route::get('appointments/patient/{patientId}', [AppointmentController::class, 'getByPatient']);
        Route::get('appointments/doctor/{doctorId}', [AppointmentController::class, 'getByDoctor']);
        Route::get('appointments/cabinet/{cabinetId}', [AppointmentController::class, 'getByCabinet']);

        // Consultation routes :doctors, patients, and admins
        Route::middleware('role:doctor,patient,admin')->group(function () {
            Route::apiResource('consultations', ConsultationController::class);
            Route::get('consultations/patient/{patientId}', [ConsultationController::class, 'getByPatient']);
            Route::get('consultations/doctor/{doctorId}', [ConsultationController::class, 'getByDoctor']);
        });

        // Prescription routes doctors, ppatients, and Admins
        Route::middleware('role:doctor,patient,admin')->group(function () {
            Route::apiResource('prescriptions', PrescriptionController::class);
            Route::get('prescriptions/patient/{patientId}', [PrescriptionController::class, 'getByPatient']);
            Route::get('prescriptions/doctor/{doctorId}', [PrescriptionController::class, 'getByDoctor']);
        });

        // Message routes  ll authenticated users
        Route::apiResource('messages', MessageController::class);
        Route::get('messages/conversation/{userId1}/{userId2}', [MessageController::class, 'getConversation']);
        Route::get('messages/user/{userId}', [MessageController::class, 'getUserMessages']);
        Route::post('messages/{id}/mark-seen', [MessageController::class, 'markAsSeen']);

        // Rating routes ^patients and admins can create, everyone also can read
        Route::get('ratings', [RatingController::class, 'index']);
        Route::get('ratings/{id}', [RatingController::class, 'show']);
        Route::get('ratings/cabinet/{cabinetId}', [RatingController::class, 'getByCabinet']);
        Route::get('ratings/patient/{patientId}', [RatingController::class, 'getByPatient']);

        Route::middleware('role:patient,admin')->group(function () {
            Route::post('ratings', [RatingController::class, 'store']);
            Route::put('ratings/{id}', [RatingController::class, 'update']);
            Route::delete('ratings/{id}', [RatingController::class, 'destroy']);
        });

        Route::middleware('role:admin')->group(function () {
            Route::post('payments', [PaymentController::class, 'store']);
            Route::put('payments/{id}', [PaymentController::class, 'update']);
            Route::delete('payments/{id}', [PaymentController::class, 'destroy']);
        });

        // "All" endpoints for authenticated users
        Route::prefix('all')->group(function () {
            Route::get('patients', [AllController::class, 'listAllPatients'])->middleware('role:doctor,admin');
            Route::get('appointments', [AllController::class, 'listAllAppointments']);
            Route::get('assistants', [AllController::class, 'listAllAssistants'])->middleware('role:admin,doctor');
            Route::get('consultations', [AllController::class, 'listAllConsultations'])->middleware('role:doctor,patient,admin');
            Route::get('users', [AllController::class, 'listAllUsers'])->middleware('role:superadmin,admin');
        });
    });
});
