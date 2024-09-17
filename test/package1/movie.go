package package1

// +import=accessor, Pkg="github.com/procyon-projects/accessor"
import "context"

// +accessor=MovieAccessor, Url="https://localhost:8090/api/v1"
type MovieAccessor interface {
	// +accessor:mapping="/movie/{id}", Method=GET
	GetMovie(ctx context.Context, id string)
	// +accessor:mapping="/movie", Method=POST
	SaveMovie(ctx context.Context)
	// +accessor:mapping="/movie/{id}", Method=PATCH
	UpdateMovie(ctx context.Context)
	// +accessor:mapping="/movie/{id}/reviews", Method=GET
	GetMovieReviews(ctx context.Context, id string)
}
