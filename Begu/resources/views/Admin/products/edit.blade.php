@extends('layouts.app')

@section('content')
    <div class="container">
        <h1>Edit Product</h1>
        <form method="POST" action="{{ route('admin.products.update', $product['ID']) }}" enctype="multipart/form-data">
            @csrf
            @method('PUT') <!-- Use PUT method for update -->
            <div class="form-group">
                <label for="category_id">Category:</label>
                <select class="form-control" id="category_id" name="category_id">
                    <option value="">Select Category</option>
                    @foreach ($categories as $category)
                        <option value="{{ $category['id'] }}" {{ $product['category_id'] == $category['id'] ? 'selected' : '' }}>
                            {{ $category['name'] }}
                        </option>
                    @endforeach
                </select>
            </div>
            <div class="form-group">
                <label for="subcategory_id">SubCategory:</label>
                <select class="form-control" id="subcategory_id" name="subcategory_id">
                    <option value="">Select SubCategory</option>
                    @foreach ($subcategories as $subcategory)
                        <option value="{{ $subcategory['id'] }}" {{ $product['subcategory_id'] == $subcategory['id'] ? 'selected' : '' }}>
                            {{ $subcategory['name'] }}
                        </option>
                    @endforeach
                </select>
            </div>
            <div class="form-group">
                <label for="brand_id">Brand:</label>
                <select class="form-control" id="brand_id" name="brand_id">
                    <option value="">Select Brand</option>
                    @foreach ($brands as $brand)
                        <option value="{{ $brand['id'] }}" {{ $product['brand_id'] == $brand['id'] ? 'selected' : '' }}>
                            {{ $brand['name'] }}
                        </option>
                    @endforeach
                </select>
            </div>
            <div class="form-group">
                <label for="name">Name:</label>
                <input type="text" class="form-control" id="name" name="name" placeholder="Enter product name" value="{{ $product['name'] }}">
            </div>
            <div class="form-group">
                <label for="size">Size:</label>
                <input type="text" class="form-control" id="size" name="size" placeholder="Enter product size" value="{{ $product['size'] }}">
            </div>
            <div class="form-group">
                <label for="quantity">Quantity:</label>
                <input type="number" class="form-control" id="quantity" name="quantity" placeholder="Enter product quantity" value="{{ $product['quantity'] }}">
            </div>
            <div class="form-group">
                <label for="price">Price:</label>
                <input type="number" step="0.01" class="form-control" id="price" name="price" placeholder="Enter product price" value="{{ $product['price'] }}">
            </div>
            <div class="form-group">
                <label for="description">Description:</label>
                <textarea class="form-control" id="description" name="description" placeholder="Enter product description">{{ $product['description'] }}</textarea>
            </div>
            <div class="form-group">
                <label for="image">Image:</label>
                @if(isset($product['image']) && Storage::disk('public')->exists($product['image']))

                <div class="mb-3">
                    <img src="{{ asset('storage/' . $product['image']) }}" alt="{{ $product['name'] }}" style="width: 100px; height: auto;">
                </div>
            @endif
            
                <input type="file" class="form-control-file" id="image" name="image">
            </div>
            <button type="submit" class="btn btn-primary">Update</button>
        </form>
    </div>
@endsection
