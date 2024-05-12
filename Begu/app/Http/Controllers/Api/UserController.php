<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class UserController extends Controller
{
    /**
     * Mengambil data user berdasarkan ID dari API Golang.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function sendUser($id)
    {
        // Mengambil instance model User berdasarkan ID
        $user = User::findOrFail($id);

        // Mengirim data user ke Golang melalui API
        $response = Http::get('http://localhost:8000/api/send-user', [
            'user' => $user->toArray(), // Mengubah model User menjadi array
        ]);

        // Menangani respons dari API
        if ($response->successful()) {
            return $response->json();
        } else {
            return response()->json(['error' => 'Failed to send user data to Golang API.'], $response->status());
        }
}


public function user($id)
{
    $user = User::findOrFail($id);

    return response()->json($user, 200);
}


}
