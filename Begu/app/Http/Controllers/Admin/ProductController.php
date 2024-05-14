<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Storage; // Import the Storage facade

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
        $response = Http::get('http://localhost:7777/api/category');
        $categories = collect($response->json())->map(function ($categories) {
            return [
                'id' => $categories['ID'],
                'name' => $categories['name']
            ];
        })->toArray();
        return view('admin.products.create', compact('subcategories','brands','categories'));
    }

    public function store(Request $request)
    {
        $request->validate([
            'name' => 'required|string|max:255',
            'subcategory_id' => 'required|integer',
            'category_id' => 'required|integer',
            'brand_id' => 'required|integer',
            'size' => 'required|string|max:255',
            'quantity' => 'required|integer',
            'price' => 'required|integer', // Ensure price is an integer
            'description' => 'required|string',
            'image' => 'required|image|mimes:jpeg,png,jpg,gif|max:2048', // Adjust the validation rules as needed
        ]);
    
        // Cast subcategory_id, category_id, and brand_id to integers
        $subcategory_id = (int) $request->input('subcategory_id');
        $category_id = (int) $request->input('category_id');
        $brand_id = (int) $request->input('brand_id');
    
        if ($request->hasFile('image')) {
            $image = $request->file('image');
            // Get the sanitized brand name
            $brandName = preg_replace('/[^a-zA-Z0-9]/', '', $request->input('name'));
            // Generate a unique filename using brand name and timestamp
            $imageName = $brandName . '_' . time() . '.' . $image->getClientOriginalExtension();
            // Store the image with the generated filename
            $imagePath = $image->storeAs('product', $imageName, 'public');
        }
    
        // Prepare the data for the HTTP request
        $data = [
            'name' => $request->input('name'),
            'subcategory_id' => $subcategory_id,
            'category_id' => $category_id,
            'brand_id' => $brand_id,
            'size' => $request->input('size'),
            'quantity' => (int) $request->input('quantity'), // Explicitly cast to integer
            'price' => (int) $request->input('price'),       // Explicitly cast to integer
                'description' => $request->input('description'),
            'image' => $imagePath ?? null,
        ];
    
        // Send the data to the external API
        $response = Http::post('http://localhost:2222/api/product', $data);
    
    
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
        $response = Http::get('http://localhost:7777/api/category');
        $categories = collect($response->json())->map(function ($categories) {
            return [
                'id' => $categories['ID'],
                'name' => $categories['name']
            ];
        })->toArray();
        return view('admin.products.edit', compact('product','subcategories','brands','categories'));
    }

    public function update(Request $request, $id)
{
    $request->validate([
        'name' => 'required|string|max:255',
        'subcategory_id' => 'required|integer',
        'category_id' => 'required|integer',
        'brand_id' => 'required|integer',
        'size' => 'required|string|max:255',
        'quantity' => 'required|integer',
        'price' => 'required|integer', // Ensure price is an integer
        'description' => 'required|string',
        'image' => 'required|image|mimes:jpeg,png,jpg,gif|max:2048', // Adjust the validation rules as needed
    ]);

    // Retrieve the existing product to get the current image path
    $existingProductResponse = Http::get("http://localhost:2222/api/product/{$id}");
    $existingProduct = $existingProductResponse->json();

    // Handle the image upload
    $imagePath = $existingProduct['image'] ?? null;
    if ($request->hasFile('image')) {
        // Delete the old image if it exists
        if ($imagePath && Storage::disk('public')->exists($imagePath)) {
            Storage::disk('public')->delete($imagePath);
        }

        $image = $request->file('image');
        $imageName = time() . '.' . $image->getClientOriginalExtension();
        $imagePath = $image->storeAs('product', $imageName, 'public');
    }


    $subcategory_id = (int) $request->input('subcategory_id');
    $category_id = (int) $request->input('category_id');
    $brand_id = (int) $request->input('brand_id');

    // Prepare the data for the HTTP request
    $data = [
        'name' => $request->input('name'),
        'subcategory_id' => $subcategory_id,
        'category_id' => $category_id,
        'brand_id' => $brand_id,
        'size' => $request->input('size'),
        'quantity' => (int) $request->input('quantity'), // Explicitly cast to integer
        'price' => (int) $request->input('price'),       // Explicitly cast to integer
    'description' => $request->input('description'),

        'image' => $request->hasFile('image') ? $imagePath : $existingProduct['image'],
    ];


    // Make the HTTP request
    $response = Http::put("http://localhost:2222/api/product/{$id}", $data);



    if ($response->successful()) {
        return redirect()->route('admin.products.index')->with('success', 'Product updated successfully.');
    } else {
        // Delete the uploaded image if the request failed
        if ($request->hasFile('image') && $imagePath) {
            Storage::disk('public')->delete($imagePath);
        }
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
