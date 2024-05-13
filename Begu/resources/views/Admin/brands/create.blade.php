<!-- resources/views/admin/categories/create.blade.php -->

@extends('layouts.app')

@section('content')
    <div class="container">
        <h1>Create Brand</h1>

        <!-- Form untuk membuat kategori baru -->
        <form method="POST" action="{{ route('admin.brands.store') }}">
            @csrf
            <div class="form-group">
                <label for="name">Name:</label>
                <input type="text" class="form-control" id="name" name="name" placeholder="Enter brand name">
            </div>
            <!-- Tambahkan field lain jika ada -->

            <button type="submit" class="btn btn-primary">Create</button>
        </form>
    </div>
@endsection
