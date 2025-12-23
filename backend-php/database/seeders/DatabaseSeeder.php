<?php

namespace Database\Seeders;

use App\Models\User;
use App\Models\Doctor;
use App\Models\Patient;
use App\Models\Cabinet;
use App\Models\Assistant;
use App\Models\Appointment;
use App\Models\Consultation;
use App\Models\Prescription;
use App\Models\Message;
use App\Models\Rating;
use App\Models\Payment;
use App\Models\SuperAdmin;
use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\Hash;

class DatabaseSeeder extends Seeder
{
    public function run(): void
    {
        // Create SuperAdmin
        $superAdminUser = User::create([
            'first_name' => 'Houria',
            'last_name' => 'Aichi',
            'email' => 'houria.aichi@medworld.dz',
            'password' => Hash::make('password123'),
            'phone_number' => '+213 555 222 000',
            'address' => 'Algiers, Algeria',
            'gender' => 'female',
            'date_of_birth' => '1985-06-15',
            'type' => 'superadmin',
        ]);
        SuperAdmin::create(['user_id' => $superAdminUser->id]);

        // Create Admin Users (Doctors who are admins)
        $adminUsers = [
            [
                'first_name' => 'Kamel',
                'last_name' => 'Daoud',
                'email' => 'kamel.daoud@medworld.dz',
                'phone_number' => '+213 555 123 456',
                'address' => 'Algiers',
                'gender' => 'male',
                'date_of_birth' => '1978-04-12',
                'speciality' => 'Family Medicine',
                'career_start' => '2002-01-01',
                'consultation_price' => 1500,
                'consultation_duration' => 30,
            ],
            [
                'first_name' => 'Yasmina',
                'last_name' => 'Khadra',
                'email' => 'yasmina.khadra@medworld.dz',
                'phone_number' => '+213 555 567 890',
                'address' => 'Algiers',
                'gender' => 'female',
                'date_of_birth' => '1984-07-08',
                'speciality' => 'Dermatology',
                'career_start' => '2008-01-01',
                'consultation_price' => 1800,
                'consultation_duration' => 25,
            ],
            [
                'first_name' => 'Mohammed',
                'last_name' => 'Dib',
                'email' => 'mohammed.dib@medworld.dz',
                'phone_number' => '+213 555 901 234',
                'address' => 'Algiers',
                'gender' => 'male',
                'date_of_birth' => '1975-03-20',
                'speciality' => 'Cardiology',
                'career_start' => '1999-01-01',
                'consultation_price' => 2500,
                'consultation_duration' => 30,
            ],
            [
                'first_name' => 'Rachid',
                'last_name' => 'Boudjedra',
                'email' => 'rachid.boudjedra@medworld.dz',
                'phone_number' => '+213 555 345 678',
                'address' => 'Algiers',
                'gender' => 'male',
                'date_of_birth' => '1982-11-02',
                'speciality' => 'Family Medicine',
                'career_start' => '2006-01-01',
                'consultation_price' => 1400,
                'consultation_duration' => 20,
            ],
            [
                'first_name' => 'Kateb',
                'last_name' => 'Yacine',
                'email' => 'kateb.yacine@medworld.dz',
                'phone_number' => '+213 555 789 012',
                'address' => 'Algiers',
                'gender' => 'male',
                'date_of_birth' => '1987-09-10',
                'speciality' => 'Dermatology',
                'career_start' => '2010-01-01',
                'consultation_price' => 1700,
                'consultation_duration' => 30,
            ],
        ];

        $adminDoctors = [];
        $cabinets = [];

        foreach ($adminUsers as $index => $adminData) {
            $user = User::create([
                'first_name' => $adminData['first_name'],
                'last_name' => $adminData['last_name'],
                'email' => $adminData['email'],
                'password' => Hash::make('password123'),
                'phone_number' => $adminData['phone_number'],
                'address' => $adminData['address'],
                'gender' => $adminData['gender'],
                'date_of_birth' => $adminData['date_of_birth'],
                'type' => 'admin',
            ]);

            // Create cabinet for admin
            $cabinet = Cabinet::create([
                'name' => 'Cabinet ' . $adminData['last_name'] . ($index + 1),
                'phone' => $adminData['phone_number'],
                'admin_id' => $user->id,
                'image' => 'https://picsum.photos/1200/700?random=' . $index,
                'access_handicap' => $index % 2 === 0,
                'has_parking' => true,
                'has_wifi' => true,
                'accepts_urgent' => $index % 2 === 0,
                'accepts_insurance' => true,
                'opening_hours' => [
                    'monday' => ['open' => '08:00', 'close' => '17:00'],
                    'tuesday' => ['open' => '08:00', 'close' => '17:00'],
                    'wednesday' => ['open' => '08:00', 'close' => '17:00'],
                    'thursday' => ['open' => '08:00', 'close' => '17:00'],
                    'friday' => ['open' => '08:00', 'close' => '12:00'],
                ],
                'location_lat' => 36.7525 + ($index * 0.01),
                'location_lng' => 3.042 + ($index * 0.01),
            ]);

            $cabinets[] = $cabinet;

            // Create doctor record for admin
            $doctor = Doctor::create([
                'user_id' => $user->id,
                'speciality' => $adminData['speciality'],
                'career_start' => $adminData['career_start'],
                'consultation_price' => $adminData['consultation_price'],
                'consultation_duration' => $adminData['consultation_duration'],
                'cabinet_id' => $cabinet->id,
            ]);

            $adminDoctors[] = $doctor;
        }

        // Create Regular Doctors
        $doctorData = [
            [
                'first_name' => 'Rayan',
                'last_name' => 'El-miloudi',
                'email' => 'rayan.miloudi@medworld.dz',
                'phone_number' => '+213 555 123 456',
                'speciality' => 'Cardiology',
                'career_start' => '2015-01-01',
                'consultation_price' => 2000,
                'consultation_duration' => 20,
            ],
            [
                'first_name' => 'Manel',
                'last_name' => 'Manel',
                'email' => 'manel.manel@medworld.dz',
                'phone_number' => '+213 555 234 567',
                'speciality' => 'Family Medicine',
                'career_start' => '2010-01-01',
                'consultation_price' => 1500,
                'consultation_duration' => 25,
            ],
            [
                'first_name' => 'Ramy',
                'last_name' => 'Hadid',
                'email' => 'ramy.ramyyy@medworld.dz',
                'phone_number' => '+213 555 345 678',
                'speciality' => 'Pediatrics',
                'career_start' => '2012-01-01',
                'consultation_price' => 1600,
                'consultation_duration' => 20,
            ],
            [
                'first_name' => 'Sarah',
                'last_name' => 'Zerrouki',
                'email' => 'sarah.zerrouki@medworld.dz',
                'phone_number' => '+213 555 456 789',
                'speciality' => 'Dermatology',
                'career_start' => '2013-01-01',
                'consultation_price' => 1800,
                'consultation_duration' => 25,
            ],
        ];

        $doctors = [];
        foreach ($doctorData as $dData) {
            $user = User::create([
                'first_name' => $dData['first_name'],
                'last_name' => $dData['last_name'],
                'email' => $dData['email'],
                'password' => Hash::make('password123'),
                'phone_number' => $dData['phone_number'],
                'type' => 'doctor',
            ]);

            $doctor = Doctor::create([
                'user_id' => $user->id,
                'speciality' => $dData['speciality'],
                'career_start' => $dData['career_start'],
                'consultation_price' => $dData['consultation_price'],
                'consultation_duration' => $dData['consultation_duration'],
                'cabinet_id' => $cabinets[0]->id,
            ]);

            $doctors[] = $doctor;
        }

        // Create Assistants
        $assistantData = [
            ['first_name' => 'Amira', 'last_name' => 'Bouraoui', 'email' => 'amira.bouraoui@medworld.dz', 'phone' => '+213 555 234 567'],
            ['first_name' => 'Assia', 'last_name' => 'Djebar', 'email' => 'assia.djebar@medworld.dz', 'phone' => '+213 555 678 901'],
            ['first_name' => 'Malika', 'last_name' => 'Mokeddem', 'email' => 'malika.mokeddem@medworld.dz', 'phone' => '+213 555 012 345'],
            ['first_name' => 'Leila', 'last_name' => 'Aissaoui', 'email' => 'leila.aissaoui@medworld.dz', 'phone' => '+213 555 456 789'],
            ['first_name' => 'Warda', 'last_name' => 'Al-Jazairia', 'email' => 'warda.aljazairia@medworld.dz', 'phone' => '+213 555 890 123'],
        ];

        $assistants = [];
        foreach ($assistantData as $index => $aData) {
            $user = User::create([
                'first_name' => $aData['first_name'],
                'last_name' => $aData['last_name'],
                'email' => $aData['email'],
                'password' => Hash::make('password123'),
                'phone_number' => $aData['phone'],
                'type' => 'assistant',
            ]);

            $assistant = Assistant::create([
                'user_id' => $user->id,
                'cabinet_id' => $cabinets[$index % count($cabinets)]->id,
            ]);

            $assistants[] = $assistant;
        }

        // Attach doctors to assistants (many to many)
        for ($i = 0; $i < count($assistants); $i++) {
            $assistants[$i]->doctors()->attach(
                array_slice(array_merge($adminDoctors, $doctors), 0, 2)
            );
        }

        // Create Patients
        $patientData = [
            [
                'first_name' => 'Riad',
                'last_name' => 'Mahrez',
                'email' => 'riad.mahrez@email.dz',
                'phone_number' => '+213 555 567 890',
                'blood_type' => 'A+',
                'medical_history' => ['Hypertension - Diagnosed 2020', 'Type 2 Diabetes'],
                'allergies' => ['Penicillin', 'Sulfa drugs'],
            ],
            [
                'first_name' => 'Amira',
                'last_name' => 'Brahimi',
                'email' => 'amira.brahimi@email.dz',
                'phone_number' => '+213 555 678 901',
                'blood_type' => 'O+',
                'medical_history' => ['Asthma - Mild', 'Migraines - Chronic'],
                'allergies' => ['Dust mites', 'Pollen', 'Aspirin'],
            ],
            [
                'first_name' => 'Karim',
                'last_name' => 'Ziani',
                'email' => 'karim.ziani@email.dz',
                'phone_number' => '+213 555 789 012',
                'blood_type' => 'B+',
                'medical_history' => ['GERD', 'Mild anxiety'],
                'allergies' => ['Shellfish', 'Ibuprofen'],
            ],
            [
                'first_name' => 'Louisa',
                'last_name' => 'Hanoune',
                'email' => 'louisa.hanoune@email.dz',
                'phone_number' => '+213 555 890 123',
                'blood_type' => 'AB+',
                'medical_history' => ['Hypothyroidism', 'Iron deficiency anemia'],
                'allergies' => ['Latex', 'Codeine'],
            ],
        ];

        $patients = [];
        foreach ($patientData as $pData) {
            $user = User::create([
                'first_name' => $pData['first_name'],
                'last_name' => $pData['last_name'],
                'email' => $pData['email'],
                'password' => Hash::make('password123'),
                'phone_number' => $pData['phone_number'],
                'type' => 'patient',
            ]);

            $patient = Patient::create([
                'user_id' => $user->id,
                'blood_type' => $pData['blood_type'],
                'medical_history' => $pData['medical_history'],
                'allergies' => $pData['allergies'],
                'emergency_contact' => '+213 555 000 000',
            ]);

            $patients[] = $patient;
        }

        // Create Appointments
        $appointmentData = [
            ['date' => '2024-01-15 09:00:00', 'status' => 'COMPLETED', 'patient' => $patients[0], 'doctor' => $adminDoctors[0], 'cabinet' => $cabinets[0]],
            ['date' => '2024-02-10 14:00:00', 'status' => 'COMPLETED', 'patient' => $patients[1], 'doctor' => $adminDoctors[1], 'cabinet' => $cabinets[1]],
            ['date' => '2024-03-05 11:00:00', 'status' => 'COMPLETED', 'patient' => $patients[2], 'doctor' => $doctors[0], 'cabinet' => $cabinets[0]],
            ['date' => '2024-04-15 09:30:00', 'status' => 'CONFIRMED', 'patient' => $patients[3], 'doctor' => $doctors[1], 'cabinet' => $cabinets[0]],
            ['date' => '2024-05-20 15:30:00', 'status' => 'CONFIRMED', 'patient' => $patients[0], 'doctor' => $adminDoctors[2], 'cabinet' => $cabinets[2]],
            ['date' => '2024-06-15 10:30:00', 'status' => 'CONFIRMED', 'patient' => $patients[1], 'doctor' => $doctors[2], 'cabinet' => $cabinets[0]],
            ['date' => '2024-07-15 10:00:00', 'status' => 'SCHEDULED', 'patient' => $patients[2], 'doctor' => $adminDoctors[3], 'cabinet' => $cabinets[3]],
        ];

        $appointments = [];
        foreach ($appointmentData as $apptData) {
            $appointment = Appointment::create([
                'date' => $apptData['date'],
                'status' => $apptData['status'],
                'patient_id' => $apptData['patient']->id,
                'doctor_id' => $apptData['doctor']->id,
                'cabinet_id' => $apptData['cabinet']->id,
                'created_by_assistant_id' => $assistants[0]->id,
            ]);
            $appointments[] = $appointment;
        }

        // Create Consultations
        $consultationData = [
            ['doctor' => $adminDoctors[0], 'patient' => $patients[0], 'appointment' => $appointments[0], 'notes' => 'Initial consultation for hypertension'],
            ['doctor' => $adminDoctors[1], 'patient' => $patients[1], 'appointment' => $appointments[1], 'notes' => 'Follow-up for asthma management'],
            ['doctor' => $doctors[0], 'patient' => $patients[2], 'appointment' => $appointments[2], 'notes' => 'GERD treatment plan'],
            ['doctor' => $doctors[1], 'patient' => $patients[3], 'appointment' => $appointments[3], 'notes' => 'Thyroid function assessment'],
        ];

        $consultations = [];
        foreach ($consultationData as $consData) {
            $consultation = Consultation::create([
                'doctor_id' => $consData['doctor']->id,
                'patient_id' => $consData['patient']->id,
                'appointment_id' => $consData['appointment']->id,
                'notes' => $consData['notes'],
                'prescriptions' => ['Medication 1', 'Medication 2'],
                'attachments' => [],
            ]);
            $consultations[] = $consultation;
        }

        // Create Prescriptions
        foreach ($consultations as $cons) {
            Prescription::create([
                'consultation_id' => $cons->id,
                'patient_id' => $cons->patient_id,
                'doctor_id' => $cons->doctor_id,
                'prescription_date' => now(),
                'status' => 'ACTIVE',
                'medications' => ['Medication A 10mg', 'Medication B 5mg'],
                'general_instructions' => 'Take with food',
                'valid_until' => now()->addMonths(3),
                'refills_allowed' => 3,
                'refills_used' => 0,
            ]);
        }

        // Create Messages
        Message::create([
            'sender_id' => $adminDoctors[0]->user_id,
            'receiver_id' => $doctors[0]->user_id,
            'cabinet_id' => $cabinets[0]->id,
            'date' => now(),
            'content' => ['message' => 'Follow-up regarding patient case'],
            'status' => 'seen',
        ]);

        Message::create([
            'sender_id' => $assistants[0]->user_id,
            'receiver_id' => $adminDoctors[1]->user_id,
            'cabinet_id' => $cabinets[1]->id,
            'date' => now(),
            'content' => ['message' => 'Appointment confirmed'],
            'status' => 'unseen',
        ]);

        // Create Ratings
        foreach ($cabinets as $cabinet) {
            Rating::create([
                'patient_id' => $patients[0]->id,
                'cabinet_id' => $cabinet->id,
                'date' => now(),
                'equippement' => ['disponibility' => true, 'rating' => 5],
                'user_experience' => ['reception' => 4, 'hygiene' => 5, 'wait_time' => 4, 'communication' => 4, 'professionalism' => 5],
                'review' => 'Great service and professional staff!',
            ]);
        }

        // Create Payments
        foreach (array_slice($appointments, 0, 3) as $appointment) {
            Payment::create([
                'patient_id' => $appointment->patient_id,
                'doctor_id' => $appointment->doctor_id,
                'cabinet_id' => $appointment->cabinet_id,
                'appointment_id' => $appointment->id,
                'amount' => 1500.00,
                'status' => 'completed',
                'payment_method' => 'cash',
                'transaction_date' => now(),
                'notes' => 'Payment received',
            ]);
        }
    }
}
