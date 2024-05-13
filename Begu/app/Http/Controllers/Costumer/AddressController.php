<?php

namespace App\Http\Controllers\Costumer;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Validator;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\Auth;
use GuzzleHttp\Client;
class AddressController extends Controller
{

    public function index()
    {
        $response = Http::get("http://localhost:9999/api/address");
        $addresses = $response->json();
        
        $userId = auth()->id(); 
        
        // Filter alamat berdasarkan user_id
        $alamat = collect($addresses)->where('user_id', $userId)->all();
    
        return view('costumer.address.index', compact('alamat'));
    }

    public function create()
    {
        return view('costumer.address.create');
    }

    public function store(Request $request)
    {
        // Validasi input jika diperlukan
        $request->validate([
            'street' => 'required|string|min:3|max:255',
            'village' => 'required|string|min:3|max:100',
            'district' => 'required|string|min:3|max:100',
            'regency' => 'required|string|min:3|max:100',
            'province' => 'required|string|min:3|max:100',
            'postal_code' => 'required|string',
            'detail' => 'required|string|min:3|max:255',
            // Tambahkan validasi untuk field lain jika ada
        ]);
    
        // Mendapatkan ID pengguna yang saat ini login
        $user_id = Auth::id();
    
        // Buat instance dari GuzzleHTTP Client
        $client = new Client();
    
        try {
            // Kirim permintaan POST ke API untuk membuat alamat baru
            $response = $client->post('http://localhost:9999/api/address', [
                'json' => [
                    'user_id' => $user_id,
                    'street' => $request->input('street'),
                    'village' => $request->input('village'),
                    'district' => $request->input('district'),
                    'regency' => $request->input('regency'),
                    'province' => $request->input('province'),
                    'postal_code' => $request->input('postal_code'),
                    'detail' => $request->input('detail'),
                    // Tambahkan field lain sesuai kebutuhan
                ]
            ]);
    
            // Periksa kode status respons
            if ($response->getStatusCode() === 201) {
                // Jika sukses, redirect ke halaman daftar alamat dengan pesan sukses
                return redirect()->route('costumer.address.index')->with('success', 'Address created successfully.');
            }
        } catch (\Exception $e) {
            // Jika terjadi kesalahan, kembalikan ke halaman pembuatan alamat dengan pesan error
            return back()->withInput()->with('error', 'Failed to create address. Please try again.');
        }
    }


}
