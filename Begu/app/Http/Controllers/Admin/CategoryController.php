<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class CategoryController extends Controller
{
    public function index()
    {
        $response = Http::get('http://localhost:7777/api/category');
        $categories = $response->json();

        return view('admin.categories.index', compact('categories'));
    }

    public function create()
    {
        return view('admin.categories.create');
    }

    public function store(Request $request)
    {
        // Validasi input jika diperlukan
        $request->validate([
            'name' => 'required|string|max:255', // Contoh validasi untuk nama kategori
            // Tambahkan validasi untuk field lain jika ada
        ]);
    
        // Kirim data ke API untuk membuat kategori baru
        $response = Http::post('http://localhost:7777/api/category', [
            'name' => $request->input('name'),
            // Tambahkan field lain sesuai kebutuhan
        ]);
    
        // Periksa jika respons dari API adalah sukses atau tidak
        if ($response->successful()) {
            // Jika sukses, redirect ke halaman daftar kategori dengan pesan sukses
            return redirect()->route('admin.categories.index')->with('success', 'Category created successfully.');
        } else {
            // Jika gagal, kembalikan ke halaman pembuatan kategori dengan pesan error
            return back()->withInput()->with('error', 'Failed to create category. Please try again.');
        }
    }
    
    public function edit($id)
    {
        $response = Http::get("http://localhost:7777/api/category/{$id}");
        $category = $response->json();

        return view('admin.categories.edit', compact('category'));
    }

    public function update(Request $request, $id)
    {
        // Validasi input jika diperlukan
        $request->validate([
            'name' => 'required|string|max:255', // Contoh validasi untuk nama kategori
            // Tambahkan validasi untuk field lain jika ada
        ]);
    
        // Kirim data ke API untuk mengupdate kategori
        $response = Http::put("http://localhost:7777/api/category/{$id}", [
            'name' => $request->input('name'),
            // Tambahkan field lain sesuai kebutuhan
        ]);

    
        // Periksa jika respons dari API adalah sukses atau tidak
        if ($response->successful()) {
            // Jika sukses, redirect ke halaman daftar kategori dengan pesan sukses
            return redirect()->route('admin.categories.index')->with('success', 'Category updated successfully.');
        } else {
            // Jika gagal, kembalikan ke halaman pembuatan kategori dengan pesan error
            return back()->withInput()->with('error', 'Failed to update category. Please try again.');
        }
    }

    public function destroy($id)
    {
        $response = Http::delete("http://localhost:7777/api/category/{$id}");
    
        // Periksa jika respons dari API adalah sukses atau tidak
        if ($response->successful()) {
            // Jika sukses, redirect ke halaman daftar kategori dengan pesan sukses
            return redirect()->route('admin.categories.index')->with('success', 'Category deleted successfully.');
        } else {
            // Jika gagal, kembalikan ke halaman daftar kategori dengan pesan error
            return back()->with('error', 'Failed to delete category. Please try again.');
        }
    }
    
}
