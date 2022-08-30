# Adapter pattern

This pattern is used when we have an old implementation and we need to add new functionality to it.
By the Open/Closed principle, which means Open for extension, Closed for modification, we cannot just modify
the "legacy" interface/struct. We should create a new interface/struct adapter so we still support "legacy" functionality
whilst adding new to it.