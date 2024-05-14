@extends('layouts.app')

@section('content')
<div class="container">
    <h1>Edit Product</h1>
    <form action="{{ route('admin.products.update', $product['id']) }}" method="POST" enctype="multipart/form-data">
        @csrf
        @method('PUT')
        <div class="form-group">
            <label for="name">SubCategory:</label>
            <select type="text" class="form-control" id="name" name="name" value="{{ $product['subcategory_id'] }}">

        </div>
        <div class="form-group">
            <label for="name">Brand:</label>
            <select type="text" class="form-control" id="name" name="name" value="{{ $product['brand_id'] }}">
        </div>
        <div class="form-group">
            <label for="name">Name:</label>
            <input type="text" class="form-control" id="name" name="name" value="{{ $product['name'] }}">
        </div>
        <div class="form-group">
            <label for="size">Size:</label>
            <input type="text" class="form-control" id="size" name="size" value="{{ $product['size'] }}">
        </div>
        <div class="form-group">
            <label for="quantity">Quantity:</label>
            <input type="text" class="form-control" id="quantity" name="quantity" value="{{ $product['quantity'] }}">
        </div>
        <div class="form-group">
            <label for="price">Price:</label>
            <input type="text" class="form-control" id="price" name="price" value="{{ $product['price'] }}">
        </div>
        <div class="form-group">
            <label for="description">Description:</label>
            <input type="text" class="form-control" id="description" name="description" value="{{ $product['description'] }}">
        </div>
        <div class="form-group">
            <label for="image">Image:</label>
            <input type="file" class="form-control-file" id="image" name="image">
        </div>
        <button type="submit" class="btn btn-primary">Update</button>
    </form>
</div>
@endsection
