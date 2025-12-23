<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasOne;

class Appointment extends Model
{
    protected $table = 'appointments';

    protected $fillable = [
        'date',
        'status',
        'patient_id',
        'doctor_id',
        'cabinet_id',
        'consultation_id',
        'created_by_assistant_id'
    ];

    protected $casts = [
        'date' => 'datetime',
    ];

    public function patient(): BelongsTo
    {
        return $this->belongsTo(Patient::class, 'patient_id');
    }

    public function doctor(): BelongsTo
    {
        return $this->belongsTo(Doctor::class, 'doctor_id');
    }

    public function cabinet(): BelongsTo
    {
        return $this->belongsTo(Cabinet::class);
    }

    public function consultation(): HasOne
    {
        return $this->hasOne(Consultation::class);
    }

    public function createdByAssistant(): BelongsTo
    {
        return $this->belongsTo(Assistant::class, 'created_by_assistant_id');
    }

    public function payment(): HasOne
    {
        return $this->hasOne(Payment::class);
    }
}
