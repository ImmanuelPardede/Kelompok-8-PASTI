<?php

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Storage; // Import the Storage facade

class PromotedController extends Controller
{
    public function index()
    {
        $response = Http::get('http://localhost:2020/api/promoted');

        if ($response->successful()) {
            $promoted = $response->json();
            return view('admin.promoted.index', compact('promoted'));
        } else {
            return back()->with('error', 'Failed to fetch promoted from API.');
        }
    }

    public function create()
    {
        return view('admin.promoted.create');
    }


    public function store(Request $request)
    {
        $request->validate([
            'title' => 'required|string|max:255',
            'image' => 'required|image|mimes:jpeg,png,jpg,gif|max:5000',
        ]);
    
       // Handle the image upload
if ($request->hasFile('image')) {
    $image = $request->file('image');
    // Get the sanitized promoted name
    $titleName = preg_replace('/[^a-zA-Z0-9]/', '', $request->input('title'));
    // Generate a unique filename using promoted name and timestamp
    $imageName = $titleName . '_' . time() . '.' . $image->getClientOriginalExtension();
    // Store the image with the generated filename
    $imagePath = $image->storeAs('promoted', $imageName, 'public');
}

    
        // Prepare the data for the HTTP request
        $data = [
            'title' => $request->input('title'),
            'image' => $imagePath ?? null, // Assuming the API accepts 'i'
        ];
    
        // Make the HTTP request
        $response = Http::post('http://localhost:2020/api/promoted', $data);
    
        if ($response->successful()) {
            return redirect()->route('admin.promoted.index')->with('success', 'promoted created successfully.');
        } else {
            // Delete the uploaded image if the request failed
            if (isset($imagePath)) {
                Storage::disk('public')->delete($imagePath);
            }
            return back()->withInput()->with('error', 'Failed to create promoted. Please try again.');
        }
    }

    public function edit($id)
    {
        $response = Http::get("http://localhost:2020/api/promoted/{$id}");
        $promoted = $response->json();

        return view('admin.promoted.edit', compact('promoted'));
    }

    public function update(Request $request, $id)
    {
        $request->validate([
            'title' => 'required|string|max:255',
            'image' => 'nullable|image|mimes:jpeg,png,jpg,gif|max:5000',
        ]);
    
        // Retrieve the existing promoted to get the current image path
        $existingpromotedResponse = Http::get("http://localhost:2020/api/promoted/{$id}");
        $existingpromoted = $existingpromotedResponse->json();
    
        // Handle the image upload
        $imagePath = $existingpromoted['image'] ?? null;
        if ($request->hasFile('image')) {
            // Delete the old image if it exists
            if ($imagePath && Storage::disk('public')->exists($imagePath)) {
                Storage::disk('public')->delete($imagePath);
            }
    
            $image = $request->file('image');
            $imageName = time() . '.' . $image->getClientOriginalExtension();
            $imagePath = $image->storeAs('promoted', $imageName, 'public');
        }
    
        // Prepare the data for the HTTP request
        $data = [
            'title' => $request->input('title'),
            // Only update the image field if a new image is uploaded
            'image' => $request->hasFile('image') ? $imagePath : $existingpromoted['image'],
        ];
    
        // Make the HTTP request
        $response = Http::put("http://localhost:2020/api/promoted/{$id}", $data);
    
        if ($response->successful()) {
            return redirect()->route('admin.promoted.index')->with('success', 'promoted updated successfully.');
        } else {
            // Delete the uploaded image if the request failed
            if ($request->hasFile('image') && $imagePath) {
                Storage::disk('public')->delete($imagePath);
            }
            return back()->withInput()->with('error', 'Failed to update promoted. Please try again.');
        }
    }
    

    public function destroy($id)
    {
        // Retrieve the existing promoted to get the image path
        $existingpromotedResponse = Http::get("http://localhost:2020/api/promoted/{$id}");
        $existingpromoted = $existingpromotedResponse->json();
    
        // Get the image path from the existing promoted data
        $imagePath = $existingpromoted['image'] ?? null;
    
        // If the image path exists and the image file exists in storage, delete it
        if ($imagePath && Storage::disk('public')->exists($imagePath)) {
            Storage::disk('public')->delete($imagePath);
        }
    
        // Make the HTTP request to delete the promoted
        $response = Http::delete("http://localhost:2020/api/promoted/{$id}");
    
        if ($response->successful()) {
            return redirect()->route('admin.promoted.index')->with('success', 'promoted deleted successfully.');
        } else {
            return back()->with('error', 'Failed to delete promoted. Please try again.');
        }
    }
}
