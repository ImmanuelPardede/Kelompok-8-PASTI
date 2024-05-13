<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class BrandController extends Controller
{
    public function index()
    {
        $response = Http::get('http://localhost:9090/api/brand');

        if ($response->successful()) {
            $brands = $response->json();

            if (!empty($brands) && is_array($brands)) {
                return view('admin.brands.index', compact('brands'));
            } else {
                // Handle empty or invalid response
                return back()->with('error', 'Empty or invalid response from API.');
            }
        } else {
            // Handle failed API request
            return back()->with('error', 'Failed to fetch brands from API.');
        }
    }


    public function create()
    {
        return view('admin.brands.create');
    }

    public function store(Request $request)
    {
        // Validasi input jika diperlukan
        $request->validate([
            'name' => 'required|string|max:255', // Contoh validasi untuk nama brands
            // Tambahkan validasi untuk field lain jika ada
        ]);

        // Kirim data ke API untuk membuat brands baru
        $response = Http::post('http://localhost:9090/api/brand', [
            'name' => $request->input('name'),
            // Tambahkan field lain sesuai kebutuhan
        ]);

        // Periksa jika respons dari API adalah sukses atau tidak
        if ($response->successful()) {
            // Jika sukses, redirect ke halaman daftar brands dengan pesan sukses
            return redirect()->route('admin.brands.index')->with('success', 'Brands created successfully.');
        } else {
            // Jika gagal, kembalikan ke halaman pembuatan brands dengan pesan error
            return back()->withInput()->with('error', 'Failed to create brands. Please try again.');
        }
    }

    public function edit($id)
    {
        $response = Http::get("http://localhost:9090/api/brand/{$id}");
        $brand = $response->json();

        return view('admin.brands.edit', compact('brand'));
    }

    public function update(Request $request, $id)
    {
        // Validasi input jika diperlukan
        $request->validate([
            'name' => 'required|string|max:255', // Contoh validasi untuk nama brands
            // Tambahkan validasi untuk field lain jika ada
        ]);

        // Kirim data ke API untuk mengupdate brands
        $response = Http::put("http://localhost:9090/api/brand/{$id}", [
            'name' => $request->input('name'),
            // Tambahkan field lain sesuai kebutuhan
        ]);


        // Periksa jika respons dari API adalah sukses atau tidak
        if ($response->successful()) {
            // Jika sukses, redirect ke halaman daftar brands dengan pesan sukses
            return redirect()->route('admin.brands.index')->with('success', 'Brands updated successfully.');
        } else {
            // Jika gagal, kembalikan ke halaman pembuatan brands dengan pesan error
            return back()->withInput()->with('error', 'Failed to update brands. Please try again.');
        }
    }

    public function destroy($id)
    {
        $response = Http::delete("http://localhost:9090/api/brand/{$id}");

        // Periksa jika respons dari API adalah sukses atau tidak
        if ($response->successful()) {
            // Jika sukses, redirect ke halaman daftar kategori dengan pesan sukses
            return redirect()->route('admin.brands.index')->with('success', 'Brands deleted successfully.');
        } else {
            // Jika gagal, kembalikan ke halaman daftar kategori dengan pesan error
            return back()->with('error', 'Failed to delete Brands. Please try again.');
        }
    }
}
