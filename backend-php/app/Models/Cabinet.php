<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Cabinet extends Model
{
    protected $table = 'cabinets';

    protected $fillable = [
        'name',
        'phone',
        'admin_id',
        'image',
        'access_handicap',
        'has_parking',
        'has_wifi',
        'accepts_urgent',
        'accepts_insurance',
        'opening_hours',
        'location_lat',
        'location_lng'
    ];

    protected $casts = [
        'access_handicap' => 'boolean',
        'has_parking' => 'boolean',
        'has_wifi' => 'boolean',
        'accepts_urgent' => 'boolean',
        'accepts_insurance' => 'boolean',
        'opening_hours' => 'array',
    ];

    public function admin(): BelongsTo
    {
        return $this->belongsTo(User::class, 'admin_id');
    }

    public function doctors(): HasMany
    {
        return $this->hasMany(Doctor::class);
    }

    public function assistants(): HasMany
    {
        return $this->hasMany(Assistant::class);
    }

    public function appointments(): HasMany
    {
        return $this->hasMany(Appointment::class);
    }

    public function consultations(): HasMany
    {
        return $this->hasMany(Consultation::class);
    }

    public function messages(): HasMany
    {
        return $this->hasMany(Message::class);
    }

    public function ratings(): HasMany
    {
        return $this->hasMany(Rating::class);
    }

    public function payments(): HasMany
    {
        return $this->hasMany(Payment::class);
    }
}
