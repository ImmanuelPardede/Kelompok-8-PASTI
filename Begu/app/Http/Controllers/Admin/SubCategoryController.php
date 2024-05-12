<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class SubCategoryController extends Controller
{
    public function index()
    {
        $response = Http::get('http://localhost:8888/api/subcategory');
        $subcategories = $response->json();

        return view('admin.subcategories.index', compact('subcategories'));
    }

    public function create()
    {
        $response = Http::get('http://localhost:7777/api/category');
        $categories = collect($response->json())->map(function ($category) {
            return [
                'id' => $category['ID'],
                'name' => $category['name']
            ];
        })->toArray();
    
        return view('admin.subcategories.create', compact('categories'));
    }
    
    

    public function store(Request $request)
    {
        // Validasi input jika diperlukan
        $request->validate([
            'name' => 'required|string|max:255',
            'category_id' => 'required',
            // Tambahkan validasi untuk field lain jika ada
        ]);
    
        $category_id = (int)$request->input('category_id');
    
        // Kirim data ke API untuk membuat kategori baru
        $response = Http::post('http://localhost:8888/api/subcategory', [
            'name' => $request->input('name'),
            'category_id' => $category_id, // Tambahkan category_id ke dalam request
            // Tambahkan field lain sesuai kebutuhan
        ]);
    
         // Periksa jika respons dari API adalah sukses atau tidak
         if ($response->successful()) {
            // Jika sukses, redirect ke halaman daftar kategori dengan pesan sukses
            return redirect()->route('admin.subcategories.index')->with('success', 'Category created successfully.');
        } else {
            // Jika gagal, kembalikan ke halaman pembuatan kategori dengan pesan error
            return back()->withInput()->with('error', 'Failed to create category. Please try again.');
        }
    }
    
    
    
    public function edit($id)
    {
        $response = Http::get("http://localhost:8888/api/subcategory/{$id}");
        $subcategory = $response->json();

        $response = Http::get('http://localhost:7777/api/category');
        $categories = collect($response->json())->map(function ($category) {
            return [
                'id' => $category['ID'],
                'name' => $category['name']
            ];
        })->toArray();

        return view('admin.subcategories.edit', compact('subcategory', 'categories'));
    }

    public function update(Request $request, $id)
    {
        // Validasi input jika diperlukan
        $request->validate([
            'name' => 'required|string|max:255',
            'category_id' => 'required',
            // Tambahkan validasi untuk field lain jika ada
        ]);
    
        $category_id = (int)$request->input('category_id');
    
        // Kirim data ke API untuk update subkategori
        $response = Http::put("http://localhost:8888/api/subcategory/{$id}", [
            'name' => $request->input('name'),
            'category_id' => $category_id, // Tambahkan category_id ke dalam request
            // Tambahkan field lain sesuai kebutuhan
        ]);
    
        // Periksa jika respons dari API adalah sukses atau tidak
        if ($response->successful()) {
            // Jika sukses, redirect ke halaman daftar kategori dengan pesan sukses
            return redirect()->route('admin.subcategories.index')->with('success', 'Category updated successfully.');
        } else {
            // Jika gagal, kembalikan ke halaman edit subkategori dengan pesan error
            return back()->withInput()->with('error', 'Failed to update category. Please try again.');
        }
    }

    public function destroy($id)
    {
        // Kirim request ke API untuk menghapus subkategori
        $response = Http::delete("http://localhost:8888/api/subcategory/{$id}");

        // Periksa jika respons dari API adalah sukses atau tidak
        if ($response->successful()) {
            // Jika sukses, redirect ke halaman daftar kategori dengan pesan sukses
            return redirect()->route('subcategories.index')->with('success', 'Category deleted successfully.');
        } else {
            // Jika gagal, kembalikan ke halaman daftar subkategori dengan pesan error
            return back()->with('error', 'Failed to delete category. Please try again.');
        }
    }
    
}
