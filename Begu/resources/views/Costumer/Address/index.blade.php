<!-- resources/views/costumer/address/index.blade.php -->

@extends('layouts.app')

@section('content')
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Addresses</div>

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
                                </tr>
                            </thead>
                            <tbody>
                                @foreach($address as $tol)
                                    <tr>
                                        <td>{{ $tol['street'] }}</td>
                                        <td>{{ $tol['village'] }}</td>
                                        <td>{{ $tol['district'] }}</td>
                                        <td>{{ $tol['regency'] }}</td>
                                        <td>{{ $tol['province'] }}</td>
                                        <td>{{ $tol['postal_code'] }}</td>
                                        <td>{{ $tol['detail'] }}</td>
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
