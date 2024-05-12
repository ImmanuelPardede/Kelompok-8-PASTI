<!-- resources/views/admin/categories/create.blade.php -->

@extends('layouts.app')

@section('content')
    <div class="container">
        <h1>Create Category</h1>

        <!-- Form untuk membuat kategori baru -->
        <form method="POST" action="{{ route('admin.categories.store') }}">
            @csrf
            <div class="form-group">
                <label for="name">Name:</label>
                <input type="text" class="form-control" id="name" name="name" placeholder="Enter category name">
            </div>
            <!-- Tambahkan field lain jika ada -->

            <button type="submit" class="btn btn-primary">Create</button>
        </form>
    </div>
@endsection
