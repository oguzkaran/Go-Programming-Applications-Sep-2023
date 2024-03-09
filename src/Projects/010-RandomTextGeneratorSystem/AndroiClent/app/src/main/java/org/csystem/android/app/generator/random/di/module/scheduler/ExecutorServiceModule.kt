package org.csystem.android.app.generator.random.di.module.scheduler

import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import java.util.concurrent.ExecutorService
import java.util.concurrent.Executors
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
object ExecutorServiceModule {
    @Provides
    @Singleton
    fun createScheduler(): ExecutorService = Executors.newCachedThreadPool()
}