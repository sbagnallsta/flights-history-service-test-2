/**
 * 
 */
package com.statravel.autoqa.repository;

import java.util.List;

import org.bson.types.ObjectId;
import org.springframework.data.mongodb.repository.MongoRepository;

import com.statravel.autoqa.domain.dto.payload.Flight;

/**
 * @author STA Development Team
 *
 */
public interface FlightsRepository extends MongoRepository<Flight, ObjectId> {
    
    /**
     * @return a list of all the flights in the history table
     */
    List<Flight> findAll();

}
