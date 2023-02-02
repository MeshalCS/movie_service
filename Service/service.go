package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "movie_service/porto/protobuf.proto"
)

type movieServer struct{

}


// Function for Latest Movies:
func (s *MovieService) LatestMovies(ctx context.Context, req *LatestMoviesRequest) (*LatestMoviesResponse, error) {
	// I will retrieve the latest movies from the The Movie DB API, when I finish the implementation code
	movies, err := getLatestMovies()
	if err != nil {
		return nil, err
	}

	// Build and return the response
	resp := &LatestMoviesResponse{
		Movies: movies,
	}
	return resp, nil
}
// Function for Search Movies:
func (s *MovieService) SearchMovies(ctx context.Context, req *SearchMoviesRequest) (*SearchMoviesResponse, error) {
	// Search the movies in the TheMovieDB APIs
	movies, err := searchMovies(req.Query)
	if err != nil {
		return nil, err
	}

	// Build and return the response
	resp := &SearchMoviesResponse{
		Movies: movies,
	}
	return resp, nil
}

// Function Add/Remove Movie from Favorites:
func (s *MovieService) AddRemoveMovieFromFavorites(ctx context.Context, req *AddRemoveMovieFromFavoritesRequest) (*AddRemoveMovieFromFavoritesResponse, error) {
	var err error
	var success bool

	// Add or remove the movie from the user's favorite list in the database
	if req.IsAdd {
		success, err = addMovieToFavorites(req.UserID, req.MovieID)
	} else {
		success, err = removeMovieFromFavorites(req.UserID, req.MovieID)
	}
	if err != nil {
		return nil, err
	}

	// Build and return the response
	resp := &AddRemoveMovieFromFavoritesResponse{
		Success: success,
	}
	return resp, nil
}

// Function for Movie Details:
func (s *MovieService) MovieDetails(ctx context.Context, req *MovieDetailsRequest) (*MovieDetailsResponse, error) {
	// I will retrieve the latest movies from the The Movie DB API, when I finish the implementation code
	movie, err := getMovieDetails(req.MovieID)
	if err != nil {
		return nil, err
	}

	// Build and return the response
	resp := &MovieDetailsResponse{
		Movie: movie,
	}
	return resp, nil
}
