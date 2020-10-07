<template>
  <div class="grid">
    <vs-row>
      <vs-col w="10" offset="2">
        <vs-input
          primary
          label-placeholder="Search movies"
          v-model="value"
          @input="retrieveMovies"
        />
      </vs-col>
    </vs-row>
    <div class="movies">
      <div v-for="movie in getMovies" :key="movie.index">
        <vs-card>
          <template #title>
            <h3>Pot with a plant</h3>
          </template>
          <template #text>
            <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit.</p>
          </template>
          <template #interactions>
            <vs-button danger icon>
              <i class="bx bx-heart"></i>
            </vs-button>
            <vs-button class="btn-chat" shadow primary>
              <i class="bx bx-chat"></i>
              <span class="span"> 54 </span>
            </vs-button>
          </template>
        </vs-card>
      </div>
    </div>
  </div>
</template>

<script>
// import { searchMoviesByTitle } from '../services/queries';

export default {
  data() {
    return {
      value: "",
      movies: [],
    };
  },
  methods: {
    retrieveMovies() {
      console.log("retrieving ", this.value);
      var el = this;
      this.searchMoviesByTitle(this.value)
        .then((results) => {
          console.log(results);
        })
        .catch((err) => console.log);
    },
    searchMoviesByTitle(title) {
      console.log("Searching movies by title");
      return this.$axios.post("/movies", {
        query: `
        query movies {
            movies(title: "${title}") {
              title
              released
              tagline
            }
          }
        `,
      });
    },
  },
  computed: {
    getMovies() {
      return this.movies;
    },
  },
};
</script>

<style>
.grid {
  width: 100%;
  height: 100vh;
  display: grid;
  /* align-items: center; */
  justify-content: center;
  padding-top: 3rem;
}

.vs-input {
  width: 30rem;
}
</style>
