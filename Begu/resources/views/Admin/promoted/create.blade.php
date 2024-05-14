@extends('layouts.app')

@section('content')
<div class="container">
    <h1>Create Promoted</h1>
    <form method="POST" action="{{ route('admin.promoted.store') }}" enctype="multipart/form-data">
        @csrf
        <div class="form-group">
            <label for="title">title:</label>
            <input type="text" class="form-control" id="title" name="title" placeholder="Enter brand title" required>
        </div>
        <div class="form-group">
            <label for="image">Image:</label>
            <input type="file" class="form-control-file" id="image" name="image" required>
        </div>
        <button type="submit" class="btn btn-primary">Create</button>
    </form>
</div>
@endsection
