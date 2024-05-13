<!-- resources/views/adminHome.blade.php -->
@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
            <div class="card">
                <div class="card-header">{{ __('Dashboard') }}</div>

                <div class="card-body">
                    You are an Admin User.
                    <br>
                    <!-- Tambahkan link menuju index categories -->
                    <a href="{{ route('admin.categories.index') }}" class="btn btn-primary mt-3">Manage Categories</a>
                    <a href="{{ route('subcategories.index') }}" class="btn btn-primary mt-3">Manage SubCategories</a>
                    <a href="{{ route('address.index')}}" class="btn btn-primary">Lihat Alamat</a>
                    <a href="{{ route('admin.brands.index')}}" class="btn btn-primary">Brands</a>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
