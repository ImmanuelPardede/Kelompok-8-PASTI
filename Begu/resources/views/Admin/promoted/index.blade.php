@extends('layouts.app')

@section('content')
<div class="container">
    <h1>Promoted</h1>
    <a href="{{ route('admin.promoted.create') }}" class="btn btn-primary mb-3">Create Promoted</a>
    <table class="table">
        <thead>
            <tr>
                <th>#</th>
                <th>title</th>
                <th>Image</th>
                <th>Action</th>
            </tr>
        </thead>
        <tbody>
            @forelse ($promoted as $ads)
            <tr>
                <td>{{ $loop->iteration }}</td>
                <td>{{ $ads['title'] }}</td> <!-- Use lowercase 'title' instead of 'title' -->
                <td>
                    @if(isset($ads['image']))
                        <img src="{{ asset('storage/' . $ads['image']) }}" alt="{{ $ads['title'] }}" style="width: 50px; height: auto;">
                    @else
                        No image available
                    @endif
                </td>
                <td>
                    <a href="{{ route('admin.promoted.edit', $ads['ID']) }}" class="btn btn-sm btn-primary">Edit</a>
                    <form action="{{ route('admin.promoted.destroy', $ads['ID']) }}" method="POST" class="d-inline">
                        @csrf
                        @method('DELETE')
                        <button type="submit" class="btn btn-sm btn-danger">Delete</button>
                    </form>
                </td>
            </tr>
            @empty
            <tr>
                <td colspan="4">No promoted found.</td>
            </tr>
            @endforelse
        </tbody>
    </table>
</div>
@endsection
