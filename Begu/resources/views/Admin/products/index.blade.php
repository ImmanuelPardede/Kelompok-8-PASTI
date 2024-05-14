@extends('layouts.app')

@section('content')
<div class="container">
    <h1>Products</h1>
    <a href="{{ route('admin.products.create') }}" class="btn btn-primary mb-3">Create Product</a>
    <table class="table">
        <thead>
            <tr>
                <th>#</th>
                <th>Name</th>
                <th>Size</th>
                <th>Quantity</th>
                <th>Price</th>
                <th>Description</th>
                <th>Image</th> <!-- New column for image -->
                <th>Action</th>
            </tr>
        </thead>
        <tbody>
            @forelse ($products as $product)
            <tr>
                <td>{{ $loop->iteration }}</td>
                <td>{{ $product['name'] }}</td>
                <td>{{ $product['size'] }}</td>
                <td>{{ $product['quantity'] }}</td>
                <td>{{ $product['price'] }}</td>
                <td>{{ $product['description'] }}</td>
                <td>
                    @if ($product['image'])
                        <img src="{{ asset('storage/' . $product['image']) }}" alt="{{ $product['name'] }}" width="100">
                    @else
                        No Image Available
                    @endif
                </td>
                <td>
                    <a href="{{ route('admin.products.edit', $product['ID']) }}" class="btn btn-sm btn-primary">Edit</a>
                    <form action="{{ route('admin.products.destroy', $product['ID']) }}" method="POST" class="d-inline">
                        @csrf
                        @method('DELETE')
                        <button type="submit" class="btn btn-sm btn-danger">Delete</button>
                    </form>
                </td>
            </tr>
            @empty
            <tr>
                <td colspan="7">No products found.</td>
            </tr>
            @endforelse
        </tbody>
    </table>
</div>
@endsection
