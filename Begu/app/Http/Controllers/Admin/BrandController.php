<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Storage; // Import the Storage facade

class BrandController extends Controller
{
    public function index()
    {
        $response = Http::get('http://localhost:9090/api/brand');

        if ($response->successful()) {
            $brands = $response->json();
            return view('admin.brands.index', compact('brands'));
        } else {
            return back()->with('error', 'Failed to fetch brands from API.');
        }
    }

    public function create()
    {
        return view('admin.brands.create');
    }


    public function store(Request $request)
    {
        $request->validate([
            'name' => 'required|string|max:255',
            'image' => 'required|image|mimes:jpeg,png,jpg,gif|max:2048', // Adjust the validation rules as needed
        ]);
    
       // Handle the image upload
if ($request->hasFile('image')) {
    $image = $request->file('image');
    // Get the sanitized brand name
    $brandName = preg_replace('/[^a-zA-Z0-9]/', '', $request->input('name'));
    // Generate a unique filename using brand name and timestamp
    $imageName = $brandName . '_' . time() . '.' . $image->getClientOriginalExtension();
    // Store the image with the generated filename
    $imagePath = $image->storeAs('brands', $imageName, 'public');
}

    
        // Prepare the data for the HTTP request
        $data = [
            'name' => $request->input('name'),
            'image' => $imagePath ?? null, // Assuming the API accepts 'i'
        ];
    
        // Make the HTTP request
        $response = Http::post('http://localhost:9090/api/brand', $data);
    
        if ($response->successful()) {
            return redirect()->route('admin.brands.index')->with('success', 'Brand created successfully.');
        } else {
            // Delete the uploaded image if the request failed
            if (isset($imagePath)) {
                Storage::disk('public')->delete($imagePath);
            }
            return back()->withInput()->with('error', 'Failed to create brand. Please try again.');
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
        $request->validate([
            'name' => 'required|string|max:255',
            'image' => 'nullable|image|mimes:jpeg,png,jpg,gif|max:2048',
        ]);
    
        // Retrieve the existing brand to get the current image path
        $existingBrandResponse = Http::get("http://localhost:9090/api/brand/{$id}");
        $existingBrand = $existingBrandResponse->json();
    
        // Handle the image upload
        $imagePath = $existingBrand['image'] ?? null;
        if ($request->hasFile('image')) {
            // Delete the old image if it exists
            if ($imagePath && Storage::disk('public')->exists($imagePath)) {
                Storage::disk('public')->delete($imagePath);
            }
    
            $image = $request->file('image');
            $imageName = time() . '.' . $image->getClientOriginalExtension();
            $imagePath = $image->storeAs('brands', $imageName, 'public');
        }
    
        // Prepare the data for the HTTP request
        $data = [
            'name' => $request->input('name'),
            // Only update the image field if a new image is uploaded
            'image' => $request->hasFile('image') ? $imagePath : $existingBrand['image'],
        ];
    
        // Make the HTTP request
        $response = Http::put("http://localhost:9090/api/brand/{$id}", $data);
    
        if ($response->successful()) {
            return redirect()->route('admin.brands.index')->with('success', 'Brand updated successfully.');
        } else {
            // Delete the uploaded image if the request failed
            if ($request->hasFile('image') && $imagePath) {
                Storage::disk('public')->delete($imagePath);
            }
            return back()->withInput()->with('error', 'Failed to update brand. Please try again.');
        }
    }
    

    public function destroy($id)
    {
        // Retrieve the existing brand to get the image path
        $existingBrandResponse = Http::get("http://localhost:9090/api/brand/{$id}");
        $existingBrand = $existingBrandResponse->json();
    
        // Get the image path from the existing brand data
        $imagePath = $existingBrand['image'] ?? null;
    
        // If the image path exists and the image file exists in storage, delete it
        if ($imagePath && Storage::disk('public')->exists($imagePath)) {
            Storage::disk('public')->delete($imagePath);
        }
    
        // Make the HTTP request to delete the brand
        $response = Http::delete("http://localhost:9090/api/brand/{$id}");
    
        if ($response->successful()) {
            return redirect()->route('admin.brands.index')->with('success', 'Brand deleted successfully.');
        } else {
            return back()->with('error', 'Failed to delete brand. Please try again.');
        }
    }
}
