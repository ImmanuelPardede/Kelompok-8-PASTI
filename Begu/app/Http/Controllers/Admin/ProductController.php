<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class ProductController extends Controller
{
    public function index()
    {
        $response = Http::get('http://localhost:2222/api/product');

        if ($response->successful()) {
            $products = $response->json();
            return view('admin.products.index', compact('products'));
        } else {
            return back()->with('error', 'Failed to fetch products from API.');
        }
    }

    public function create()
    {
        $response = Http::get('http://localhost:8888/api/subcategory');
        $subcategories = collect($response->json())->map(function ($subcategory) {
            return [
                'id' => $subcategory['ID'],
                'name' => $subcategory['name']
            ];
        })->toArray();
        $response = Http::get('http://localhost:9090/api/brand');
        $brands = collect($response->json())->map(function ($brand) {
            return [
                'id' => $brand['ID'],
                'name' => $brand['name']
            ];
        })->toArray();
        return view('admin.products.create', compact('subcategories','brands'));
    }

    public function store(Request $request)
    {
        $request->validate([
            'name' => 'required|string|max:255',
            'subcategory_id' => 'required|integer',
            'brand_id' => 'required|integer',
            'size' => 'required|string|max:255',
            'quantity' => 'required|integer',
            'price' => 'required|numeric',
            'description' => 'required|string',
            'image' => 'required|file|image|max:10240', // max 10MB
        ]);

        $response = Http::post('http://localhost:2222/api/product', $request->all());

        if ($response->successful()) {
            return redirect()->route('admin.products.index')->with('success', 'Product created successfully.');
        } else {
            return back()->withInput()->with('error', 'Failed to create product. Please try again.');
        }
    }

    public function edit($id)
    {
        $response = Http::get("http://localhost:2222/api/product/{$id}");
        $product = $response->json();

        return view('admin.products.edit', compact('product'));
    }

    public function update(Request $request, $id)
    {
        $request->validate([
            'name' => 'required|string|max:255',
            'subcategory_id' => 'required|integer',
            'brand_id' => 'required|integer',
            'size' => 'required|string|max:255',
            'quantity' => 'required|integer',
            'price' => 'required|numeric',
            'description' => 'required|string',
            'image' => 'nullable|file|image|max:10240', // max 10MB
        ]);

        $response = Http::put("http://localhost:2222/api/product/{$id}", $request->all());

        if ($response->successful()) {
            return redirect()->route('admin.products.index')->with('success', 'Product updated successfully.');
        } else {
            return back()->withInput()->with('error', 'Failed to update product. Please try again.');
        }
    }

    public function destroy($id)
    {
        $response = Http::delete("http://localhost:2222/api/product/{$id}");

        if ($response->successful()) {
            return redirect()->route('admin.products.index')->with('success', 'Product deleted successfully.');
        } else {
            return back()->with('error', 'Failed to delete product. Please try again.');
        }
    }
}
