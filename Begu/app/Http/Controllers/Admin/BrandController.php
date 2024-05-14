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
        ]);

        $response = Http::post('http://localhost:9090/api/brand', [
            'name' => $request->input('name'),
        ]);

        if ($response->successful()) {
            return redirect()->route('admin.brands.index')->with('success', 'Brand created successfully.');
        } else {
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
        ]);

        $response = Http::put("http://localhost:9090/api/brand/{$id}", [
            'name' => $request->input('name'),
        ]);

        if ($response->successful()) {
            return redirect()->route('admin.brands.index')->with('success', 'Brand updated successfully.');
        } else {
            return back()->withInput()->with('error', 'Failed to update brand. Please try again.');
        }
    }

    public function destroy($id)
    {
        $response = Http::delete("http://localhost:9090/api/brand/{$id}");

        if ($response->successful()) {
            return redirect()->route('admin.brands.index')->with('success', 'Brand deleted successfully.');
        } else {
            return back()->with('error', 'Failed to delete brand. Please try again.');
        }
    }
}
