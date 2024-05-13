<!-- resources/views/admin/brands/index.blade.php -->

@extends('layouts.app')

@section('content')
    <div class="container">
        <h1>Brands</h1>

        <!-- Tombol untuk membuat brands baru -->
        <a href="{{ route('admin.brands.create') }}" class="btn btn-primary mb-3">Create Brand</a>

        <!-- Tabel untuk menampilkan daftar brands -->
        <table class="table">
            <thead>
                <tr>
                    <th>#</th>
                    <th>Name</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                @forelse ($brands as $brand)
                    <tr>
                        <td>{{ $loop->iteration }}</td> <!-- Nomor urut -->
                        <td>{{ $brand['name'] }}</td> <!-- Nama brands -->
                        <td>
                            <!-- Tombol edit -->
                            <a href="{{ route('admin.brands.edit', $brand['ID']) }}" class="btn btn-sm btn-primary">Edit</a>

                            <!-- Form untuk menghapus brands -->
                            <form action="{{ route('admin.brands.destroy', $brand['ID']) }}" method="POST" class="d-inline">
                                @csrf
                                @method('DELETE')
                                <button type="submit" class="btn btn-sm btn-danger">Delete</button>
                            </form>
                        </td>
                    </tr>
                @empty
                    <tr>
                        <td colspan="3">No brands found.</td>
                    </tr>
                @endforelse
            </tbody>
        </table>
    </div>
@endsection
