<?php
require 'vendor/autoload.php';

$app = require_once 'bootstrap/app.php';
$app->make('Illuminate\Contracts\Console\Kernel')->bootstrap();

$user = \App\Models\User::where('email', 'houria.aichi@medworld.dz')->first();
if ($user) {
    $testPassword = 'password123';
    $match = \Illuminate\Support\Facades\Hash::check($testPassword, $user->password);
    echo "Password test result: " . ($match ? "MATCH" : "NO MATCH") . "\n";
    echo "User ID: " . $user->id . "\n";
    echo "Email: " . $user->email . "\n";
} else {
    echo "User not found\n";
}
