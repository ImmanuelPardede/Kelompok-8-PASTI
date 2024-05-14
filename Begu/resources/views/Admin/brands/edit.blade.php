@extends('layouts.app')

@section('content')
<div class="container">
    <h1>Edit Brand</h1>
    <form action="{{ route('admin.brands.update', $brand['ID']) }}" method="POST" enctype="multipart/form-data">
        @csrf
        @method('PUT')
        <div class="form-group">
            <label for="name">Brand Name:</label>
            <input type="text" class="form-control" id="name" name="name" value="{{ $brand['name'] }}" required>
        </div>
        <div class="form-group">
            <label for="image">Brand Image:</label>
            @if(isset($brand['image']) && Storage::disk('public')->exists($brand['image']))
                <div class="mb-3">
                    <img src="{{ asset('storage/' . $brand['image']) }}" alt="{{ $brand['name'] }}" style="width: 100px; height: auto;">
                </div>
            @endif
            <input type="file" class="form-control-file" id="image" name="image">
        </div>
        <button type="submit" class="btn btn-primary">Update</button>
    </form>
</div>
@endsection
