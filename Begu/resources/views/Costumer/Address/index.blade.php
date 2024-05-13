<!-- resources/views/costumer/address/index.blade.php -->

@extends('layouts.app')

@section('content')
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Addresses</div>
                    <a href="{{ route('address.create') }}" class="btn btn-primary">Create Address</a>

                    <div class="card-body">
                        <table class="table">
                            <thead>
                                <tr>
                                    <th>Street</th>
                                    <th>Village</th>
                                    <th>District</th>
                                    <th>Regency</th>
                                    <th>Province</th>
                                    <th>Postal Code</th>
                                    <th>Detail</th>
                                    <th>user</th>
                                    <th>Actions</th> <!-- Tambah kolom Actions -->
                                </tr>
                            </thead>
                            <tbody>
                                @foreach($alamat as $tol)
                                    <tr>
                                        <td>{{ $tol['street'] }}</td>
                                        <td>{{ $tol['village'] }}</td>
                                        <td>{{ $tol['district'] }}</td>
                                        <td>{{ $tol['regency'] }}</td>
                                        <td>{{ $tol['province'] }}</td>
                                        <td>{{ $tol['postal_code'] }}</td>
                                        <td>{{ $tol['detail'] }}</td>
                                        <td>{{ $tol['user_id'] }}</td>
                                        <td>
                                            <a href="{{ route('address.edit', $tol['ID']) }}" class="btn btn-sm btn-primary">Edit</a> <!-- Tambah tombol Edit -->
                                        </td>
                                    </tr>
                                @endforeach
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection
