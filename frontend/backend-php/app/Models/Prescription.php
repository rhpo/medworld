<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Prescription extends Model
{
    protected $table = 'prescriptions';

    protected $fillable = [
        'consultation_id',
        'patient_id',
        'doctor_id',
        'prescription_date',
        'status',
        'medications',
        'general_instructions',
        'valid_until',
        'refills_allowed',
        'refills_used'
    ];

    protected $casts = [
        'prescription_date' => 'datetime',
        'medications' => 'array',
        'valid_until' => 'date',
    ];

    public function consultation(): BelongsTo
    {
        return $this->belongsTo(Consultation::class);
    }

    public function patient(): BelongsTo
    {
        return $this->belongsTo(Patient::class);
    }

    public function doctor(): BelongsTo
    {
        return $this->belongsTo(Doctor::class);
    }
}
