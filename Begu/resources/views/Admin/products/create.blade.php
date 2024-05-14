@extends('layouts.app')

@section('content')
    <div class="container">
        <h1>Create Product</h1>
        <form method="POST" action="{{ route('admin.products.store') }}" enctype="multipart/form-data">
            @csrf
            <div class="form-group">
                <label for="subcategory_id">SubCategory:</label>
                <select class="form-control" id="subcategory_id" name="subcategory_id" placeholder="Select SubCategory">
                    <option value="">Select SubCategory</option>
                    @foreach ($subcategories as $subcategory)
                        <option value="{{ $subcategory['id'] }}">{{ $subcategory['name'] }}</option>
                    @endforeach
                </select>
            </div>
            <div class="form-group">
                <label for="brand_id">Brand:</label>
                <select class="form-control" id="brand_id" name="brand_id" placeholder="Select Brand">
                    <option value="">Select Category</option>
                    @foreach ($brands as $brand)
                        <option value="{{ $brand['id'] }}">{{ $brand['name'] }}</option>
                    @endforeach
                </select>
            </div>
            <div class="form-group">
                <label for="name">Name:</label>
                <input type="text" class="form-control" id="name" name="name" placeholder="Enter product name">
            </div>
            <div class="form-group">
                <label for="size">Size:</label>
                <input type="text" class="form-control" id="size" name="size" placeholder="Enter product size">
            </div>
            <div class="form-group">
                <label for="quantity">Quantity:</label>
                <input type="text" class="form-control" id="quantity" name="quantity"
                    placeholder="Enter product quantity">
            </div>
            <div class="form-group">
                <label for="price">Price:</label>
                <input type="text" class="form-control" id="price" name="price" placeholder="Enter product price">
            </div>
            <div class="form-group">
                <label for="description">Description:</label>
                <input type="text" class="form-control" id="description" name="description"
                    placeholder="Enter product description">
            </div>
            <div class="form-group">
                <label for="image">Image:</label>
                <input type="file" class="form-control-file" id="image" name="image">
            </div>
            <button type="submit" class="btn btn-primary">Create</button>
        </form>
    </div>
@endsection
