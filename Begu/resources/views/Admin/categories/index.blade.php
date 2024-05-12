<!-- resources/views/admin/categories/index.blade.php -->

@extends('layouts.app')

@section('content')
    <div class="container">
        <h1>Categories</h1>

        <!-- Tombol untuk membuat kategori baru -->
        <a href="{{ route('admin.categories.create') }}" class="btn btn-primary mb-3">Create Category</a>

        <!-- Tabel untuk menampilkan daftar kategori -->
        <table class="table">
            <thead>
                <tr>
                    <th>#</th>
                    <th>Name</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                @forelse ($categories as $category)
                    <tr>
                        <td>{{ $loop->iteration }}</td> <!-- Nomor urut -->
                        <td>{{ $category['name'] }}</td> <!-- Nama kategori -->
                        <td>
                            <!-- Tombol edit -->
                            <a href="{{ route('admin.categories.edit', $category['ID']) }}" class="btn btn-sm btn-primary">Edit</a>

                            <!-- Form untuk menghapus kategori -->
                            <form action="{{ route('admin.categories.destroy', $category['ID']) }}" method="POST" class="d-inline">
                                @csrf
                                @method('DELETE')
                                <button type="submit" class="btn btn-sm btn-danger">Delete</button>
                            </form> 
                        </td>
                    </tr>
                @empty
                    <tr>
                        <td colspan="3">No categories found.</td>
                    </tr>
                @endforelse
            </tbody>
        </table>
    </div>
@endsection
