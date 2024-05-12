@extends('layouts.app')

@section('content')
    <div class="container">
        <div class="row">
            <div class="col-md-12">
                <h1>Subcategories</h1>
                <a href="{{ route('subcategories.create') }}" class="btn btn-primary">Create Subcategory</a>
                <table class="table mt-3">
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Category ID</th>
                            <th>Actions</th> <!-- New column for actions -->
                        </tr>
                    </thead>
                    <tbody>
                        @foreach ($subcategories as $subcategory)
                            <tr>
                                <td>{{ $subcategory['name'] }}</td>
                                <td>{{ $subcategory['category_id'] }}</td>
                                <td>
                                    <a href="{{ route('subcategories.edit', $subcategory['ID']) }}" class="btn btn-primary">Edit</a>
                                    <form action="{{ route('subcategories.destroy', $subcategory['ID']) }}" method="POST" style="display: inline-block;">
                                        @csrf
                                        @method('DELETE')
                                        <button type="submit" class="btn btn-danger">Delete</button>
                                    </form>
                                </td>
                            </tr>
                        @endforeach
                    </tbody>
                </table>
            </div>
        </div>
    </div>
@endsection
