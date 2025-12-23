<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Consultation extends Model
{
    protected $table = 'consultations';

    protected $fillable = [
        'doctor_id',
        'patient_id',
        'appointment_id',
        'notes',
        'prescriptions',
        'attachments'
    ];

    protected $casts = [
        'prescriptions' => 'array',
        'attachments' => 'array',
    ];

    public function doctor(): BelongsTo
    {
        return $this->belongsTo(Doctor::class);
    }

    public function patient(): BelongsTo
    {
        return $this->belongsTo(Patient::class);
    }

    public function appointment(): BelongsTo
    {
        return $this->belongsTo(Appointment::class);
    }
}
