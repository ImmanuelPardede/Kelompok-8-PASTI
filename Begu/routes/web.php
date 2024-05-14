<?php

use Illuminate\Support\Facades\Route;
use Illuminate\Support\Facades\Auth;
use App\Http\Controllers\Admin\SubcategoryController;
use App\Http\Controllers\Admin\BrandController;
use App\Http\Controllers\Admin\PromotedController;
use App\Http\Controllers\HomeController;
use App\Http\Controllers\Admin\CategoryController;
use App\Http\Controllers\Costumer\AddressController;

Route::get('/', function () {
    return view('welcome');
});

Auth::routes();

//Normal Users Routes List
Route::middleware(['auth', 'user-access:user'])->group(function () {

    Route::get('/home', [HomeController::class, 'index'])->name('home');
});

//Admin Routes List
Route::middleware(['auth', 'user-access:admin'])->group(function () {

    Route::get('/admin/home', [HomeController::class, 'adminHome'])->name('admin.home');
    Route::get('/categories', [CategoryController::class, 'index'])->name('admin.categories.index');
    Route::get('/categories/create', [CategoryController::class, 'create'])->name('admin.categories.create');
    Route::post('categories', [CategoryController::class, 'store'])->name('admin.categories.store');
    Route::get('/categories/{id}/edit', [CategoryController::class, 'edit'])->name('admin.categories.edit');
    Route::put('/categories/{id}', [CategoryController::class, 'update'])->name('admin.categories.update');
    Route::delete('/admin/categories/{id}', [CategoryController::class, 'destroy'])->name('admin.categories.destroy');
    Route::resource('subcategories', SubCategoryController::class);

    Route::resource('brands', BrandController::class)->except(['show']);
    Route::get('/brands', [BrandController::class, 'index'])->name('admin.brands.index');
    Route::get('/brands/create', [BrandController::class, 'create'])->name('admin.brands.create');
    Route::post('brands', [BrandController::class, 'store'])->name('admin.brands.store');
    Route::get('/brands/{id}/edit', [BrandController::class, 'edit'])->name('admin.brands.edit');
    Route::put('/brands/{id}', [BrandController::class, 'update'])->name('admin.brands.update');
    Route::delete('/admin/brands/{id}', [BrandController::class, 'destroy'])->name('admin.brands.destroy');

    Route::get('/promoted', [PromotedController::class, 'index'])->name('admin.promoted.index');
    Route::get('/promoted/create', [PromotedController::class, 'create'])->name('admin.promoted.create');
    Route::post('promoted', [PromotedController::class, 'store'])->name('admin.promoted.store');
    Route::get('/promoted/{id}/edit', [PromotedController::class, 'edit'])->name('admin.promoted.edit');
    Route::put('/promoted/{id}', [PromotedController::class, 'update'])->name('admin.promoted.update');
    Route::delete('/admin/promoted/{id}', [PromotedController::class, 'destroy'])->name('admin.promoted.destroy');

});

Route::get('/address', [AddressController::class, 'index'])->name('address.index');
Route::get('/addresses/create', [AddressController::class, 'create'])->name('address.create');
Route::post('/addresses', [AddressController::class, 'store'])->name('address.store');
Route::get('/address/{id}/edit', [AddressController::class, 'edit'])->name('address.edit');
Route::put('/address/{id}', [AddressController::class, 'update'])->name('address.update');
