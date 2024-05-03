@extends('layouts.management.master')

@section('content')

<div class="col-lg-12 grid-margin stretch-card">
    <div class="card">
        <div class="card-body">
            <div class="d-flex justify-content-between align-items-center mb-3">
                <h1 class="card-title">List Berita</h1>
                <!-- Tampilkan notifikasi jika ada -->
                @if (session('success'))
                    <div class="alert alert-success">
                        {{ session('success') }}
                    </div>
                @endif
                <a href="{{ route('berita.create') }}" class="btn btn-success mb-3">Tambahkan Berita</a>
            </div>

            <div class="table-responsive">
                <table class="table mt-3">
                    <thead>
                        <tr>
                            <th>No</th>
                            <th>Image</th>
                            <th>Judul</th>
                            <th>Kategori</th>
                            <th>Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        @foreach($berita as $index => $beritaItem)
                        <tr>
                            <td>{{ $index +1}}</td>
                            <td style="max-width: 200px;">
                                <img src="{{ asset($beritaItem->img_berita) }}" alt="berita Image" class="img-fluid" style="border-radius: initial; width: 100%; height: auto; max-width: 100%;">
                            </td>
                            <td>{{$beritaItem->judul }}</td>
                            <td>{{$beritaItem->kategori->kategori }}</td>
                            <td>
                                <a href="{{ route('berita.show', $beritaItem->id) }}" class="btn btn-primary btn-sm">Detail</a>
                                <a href="{{ route('berita.edit', $beritaItem->id) }}" class="btn btn-info btn-sm">Edit</a>
                                <form action="{{ route('berita.destroy', $beritaItem->id) }}" method="POST" style="display: inline-block;">
                                    @csrf
                                    @method('DELETE')
                                    <button type="submit" class="btn btn-danger btn-sm" onclick="return confirm('Are you sure?')">Delete</button>
                                </form>
                            </td>
                        </tr>
                        @endforeach
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

@endsection