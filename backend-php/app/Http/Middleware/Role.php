<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class Role
{
    /**
     * Handle an incoming request.
     */
    public function handle(Request $request, Closure $next, ...$roles): Response
    {
        // Get the authenticated user
        $user = $request->user();

        // If no user is authenticated, throw exception
        if (!$user) {
            throw new \Illuminate\Auth\AuthenticationException('Unauthenticated');
        }

        // Check if user's type is in the allowed roles
        if (in_array($user->type, $roles)) {
            return $next($request);
        }

        // User doesn't have the required role
        return response()->json([
            'message' => 'Unauthorized. Required role: ' . implode(', ', $roles),
            'user_role' => $user->type,
        ], 403);
    }
}
