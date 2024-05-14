@extends('layouts.app')

@section('content')
<div class="container">
    <h1>Edit Brand</h1>
    <form action="{{ route('admin.promoted.update', $promoted['ID']) }}" method="POST" enctype="multipart/form-data">
        @csrf
        @method('PUT')
        <div class="form-group">
            <label for="title">title:</label>
            <input type="text" class="form-control" id="title" name="title" value="{{ $promoted['title'] }}" required>
        </div>
        <div class="form-group">
            <label for="image">promoted Image:</label>
            @if(isset($promoted['image']) && Storage::disk('public')->exists($promoted['image']))
                <div class="mb-3">
                    <img src="{{ asset('storage/' . $promoted['image']) }}" alt="{{ $promoted['title'] }}" style="width: 100px; height: auto;">
                </div>
            @endif
            <input type="file" class="form-control-file" id="image" name="image">
        </div>
        <button type="submit" class="btn btn-primary">Update</button>
    </form>
</div>
@endsection
