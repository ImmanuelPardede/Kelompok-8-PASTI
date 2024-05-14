@extends('layouts.app')

@section('content')
<div class="container">
    <h1>Brands</h1>
    <a href="{{ route('admin.brands.create') }}" class="btn btn-primary mb-3">Create Brand</a>
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
                <td>{{ $loop->iteration }}</td>
                <td>{{ $brand['name'] }}</td>
                <td>
                    <a href="{{ route('admin.brands.edit', $brand['ID']) }}" class="btn btn-sm btn-primary">Edit</a>
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
