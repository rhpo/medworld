<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Rating extends Model
{
    protected $table = 'ratings';

    protected $fillable = [
        'patient_id',
        'cabinet_id',
        'date',
        'equippement',
        'user_experience',
        'review'
    ];

    protected $casts = [
        'date' => 'datetime',
        'equippement' => 'array',
        'user_experience' => 'array',
    ];

    public function patient(): BelongsTo
    {
        return $this->belongsTo(Patient::class);
    }

    public function cabinet(): BelongsTo
    {
        return $this->belongsTo(Cabinet::class);
    }
}
