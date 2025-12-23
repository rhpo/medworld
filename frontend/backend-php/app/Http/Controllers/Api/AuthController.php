<?php

namespace App\Http\Controllers\Api;

use App\Models\User;
use App\Http\Requests\LoginRequest;
use App\Http\Requests\RegisterRequest;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;
use Symfony\Component\HttpFoundation\Cookie;

class AuthController
{
    public function login(LoginRequest $request)
    {
        try {
            \Log::debug('Login attempt', ['email' => $request->email, 'method' => $request->method()]);

            $validated = $request->validated();
            \Log::debug('Validated data', $validated);

            $user = User::where('email', $validated['email'])->first();

            if (!$user) {
                \Log::debug('User not found', ['email' => $validated['email']]);
                return response()->json(['message' => 'User not found'], 401);
            }

            if (!Hash::check($validated['password'], $user->password)) {
                \Log::debug('Password mismatch', ['email' => $validated['email'], 'password' => $validated['password']]);
                return response()->json(['message' => 'Invalid password'], 401);
            }

            \Log::debug('Login successful', ['user_id' => $user->id]);
            $token = $user->createToken('api-token')->plainTextToken;

            // Create the response with the token
            $response = response()->json([
                'message' => 'Login successful',
                'user' => $user->load(['doctor', 'patient']),
                'token' => $token,
            ], 200);

            // Add the auth token cookie
            $cookie = new Cookie(
                name: 'auth_token',
                value: $token,
                expire: time() + (60 * 60 * 24 * 7), // 7 days
                path: '/',
                domain: null,
                secure: false,
                httpOnly: true,
                raw: false,
                sameSite: 'lax'
            );

            return $response->withCookie($cookie);
        } catch (\Exception $e) {
            \Log::error('Login error', ['error' => $e->getMessage(), 'trace' => $e->getTraceAsString()]);
            return response()->json(['message' => 'Server error', 'error' => $e->getMessage()], 500);
        }
    }

    public function register(RegisterRequest $request)
    {
        $validated = $request->validated();

        $user = User::create([
            'first_name' => $validated['first_name'],
            'last_name' => $validated['last_name'],
            'email' => $validated['email'],
            'password' => Hash::make($validated['password']),
            'phone_number' => $validated['phone_number'] ?? null,
            'address' => $validated['address'] ?? null,
            'gender' => $validated['gender'] ?? null,
            'date_of_birth' => $validated['date_of_birth'] ?? null,
            'type' => $validated['type'],
        ]);

        $token = $user->createToken('api-token')->plainTextToken;

        // Create the response with the token
        $response = response()->json([
            'message' => 'Registration successful',
            'user' => $user->load(['doctor', 'patient']),
            'token' => $token,
        ], 201);

        // Add the auth token cookie
        $cookie = new Cookie(
            name: 'auth_token',
            value: $token,
            expire: time() + (60 * 60 * 24 * 7), // 7 days
            path: '/',
            domain: null,
            secure: false,
            httpOnly: true,
            raw: false,
            sameSite: 'lax'
        );

        return $response->withCookie($cookie);
    }

    public function logout(Request $request)
    {
        $request->user()->currentAccessToken()->delete();

        return response()->json(['message' => 'Logout successful'], 200);
    }

    public function me(Request $request)
    {
        return response()->json([
            'user' => $request->user()->load(['doctor', 'patient']),
        ], 200);
    }
}
