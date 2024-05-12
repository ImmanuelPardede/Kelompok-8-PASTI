<?php

namespace App\Http\Controllers\Costumer;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class AddressController extends Controller
{
    public function index()
    {
        
        $response = Http::get('http://localhost:9999/api/address');
        $address = $response->json();

        return view('costumer.address.index', compact('address'));
    }
}
