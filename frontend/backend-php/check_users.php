<?php
require 'vendor/autoload.php';

$app = require_once 'bootstrap/app.php';
$app->make('Illuminate\Contracts\Console\Kernel')->bootstrap();

$users = \App\Models\User::all();
echo "Total users: " . count($users) . "\n";
foreach ($users as $user) {
    echo "- " . $user->email . "\n";
}

$target = \App\Models\User::where('email', 'houria.aichi@medworld.dz')->first();
if ($target) {
    echo "\nTarget user found: " . $target->email . "\n";
    echo "Has password hash: " . (!empty($target->password) ? "YES" : "NO") . "\n";
} else {
    echo "\nTarget user NOT found\n";
}
