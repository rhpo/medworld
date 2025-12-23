<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class AuthenticateFromCookie
{
    /**
     * Handle an incoming request.
     */
    public function handle(Request $request, Closure $next): Response
    {
        // Log for debugging
        \Log::debug('AuthenticateFromCookie middleware called', [
            'path' => $request->path(),
            'bearer_token' => $request->bearerToken() ? 'yes' : 'no',
            'auth_token_cookie' => $request->cookie('auth_token') ? 'yes' : 'no',
        ]);

        // If there's already a Bearer token in the Authorization header, let it pass through
        if ($request->bearerToken()) {
            \Log::debug('Bearer token found, passing through');
            return $next($request);
        }

        // Check if token exists in cookie
        if ($token = $request->cookie('auth_token')) {
            \Log::debug('auth_token cookie found, setting Authorization header');
            // Add the token to the Authorization header so Sanctum can authenticate it
            $request->headers->set('Authorization', 'Bearer ' . $token);
        } else {
            \Log::debug('No auth_token cookie found');
        }

        return $next($request);
    }
}
