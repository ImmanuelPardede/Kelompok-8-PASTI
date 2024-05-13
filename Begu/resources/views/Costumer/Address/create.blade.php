@extends('layouts.app')

@section('content')
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Create Address</div>
                    <div class="card-body">
                        @if ($errors->any())
                            <div class="alert alert-danger">
                                <ul>
                                    @foreach ($errors->all() as $error)
                                        <li>{{ $error }}</li>
                                    @endforeach
                                </ul>
                            </div>
                        @endif
                        <form method="POST" action="{{ route('address.store') }}">
                            @csrf
                            <div class="form-group">
                                <label for="street">Street</label>
                                <input type="text" class="form-control" id="street" name="street" value="{{ old('street') }}">
                            </div>
                            <div class="form-group">
                                <label for="village">Village</label>
                                <input type="text" class="form-control" id="village" name="village" value="{{ old('village') }}">
                            </div>
                            <div class="form-group">
                                <label for="district">District</label>
                                <input type="text" class="form-control" id="district" name="district" value="{{ old('district') }}">
                            </div>
                            <div class="form-group">
                                <label for="regency">Regency</label>
                                <input type="text" class="form-control" id="regency" name="regency" value="{{ old('regency') }}">
                            </div>
                            <div class="form-group">
                                <label for="province">Province</label>
                                <input type="text" class="form-control" id="province" name="province" value="{{ old('province') }}">
                            </div>
                            <div class="form-group">
                                <label for="postal_code">Postal Code</label>
                                <input type="text" class="form-control" id="postal_code" name="postal_code" value="{{ old('postal_code') }}">
                            </div>
                            <div class="form-group">
                                <label for="detail">Detail</label>
                                <textarea class="form-control" id="detail" name="detail">{{ old('detail') }}</textarea>
                            </div>
                            <button type="submit" class="btn btn-primary">Submit</button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection
