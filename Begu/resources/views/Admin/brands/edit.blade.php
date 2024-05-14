@extends('layouts.app')

@section('content')
<div class="container">
    <h1>Edit Brand</h1>
    <form action="{{ route('admin.brands.update', $brand['ID']) }}" method="POST">
        @csrf
        @method('PUT')
        <div class="form-group">
            <label for="name">Brand Name:</label>
            <input type="text" class="form-control" id="name" name="name" value="{{ $brand['name'] }}">
        </div>
        <button type="submit" class="btn btn-primary">Update</button>
    </form>
</div>
@endsection
