package org.csystem.android.app.generator.random;

import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.os.Message;
import android.util.Log;
import android.widget.ArrayAdapter;
import android.widget.Toast;
import androidx.appcompat.app.AppCompatActivity;
import androidx.databinding.DataBindingUtil;
import dagger.hilt.android.AndroidEntryPoint;
import org.csystem.android.app.generator.random.databinding.ActivityMainBinding;
import org.csystem.android.app.generator.random.viewmodel.data.GenerateInfo;
import org.csystem.android.app.generator.random.viewmodel.data.ServerInfo;
import org.csystem.android.app.generator.random.viewmodel.listener.MainActivityListenerViewModel;
import java.io.BufferedReader;
import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.lang.ref.WeakReference;
import java.net.Socket;
import java.nio.charset.StandardCharsets;
import java.util.concurrent.ExecutorService;
import javax.inject.Inject;

@dagger.hilt.android.AndroidEntryPoint
@kotlin.Metadata(mv = {1, 9, 0}, k = 1, xi = 48, d1 = {"\u0000@\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0005\n\u0002\u0010\u0002\n\u0002\b\u0004\n\u0002\u0018\u0002\n\u0002\b\u0004\n\u0002\u0010\u000e\n\u0002\b\t\n\u0002\u0018\u0002\n\u0002\b\u0004\b\u0007\u0018\u00002\u00020\u0001:\u0001%B\u0005\u00a2\u0006\u0002\u0010\u0002J\u0006\u0010\r\u001a\u00020\u000eJ\u0006\u0010\u000f\u001a\u00020\u000eJ\b\u0010\u0010\u001a\u00020\u000eH\u0002J\u0010\u0010\u0011\u001a\u00020\u000e2\u0006\u0010\u0012\u001a\u00020\u0013H\u0002J\u0006\u0010\u0014\u001a\u00020\u000eJ\u0010\u0010\u0015\u001a\u00020\u000e2\u0006\u0010\u0012\u001a\u00020\u0013H\u0002J\u0010\u0010\u0016\u001a\u00020\u000e2\u0006\u0010\u0017\u001a\u00020\u0018H\u0002J\u0010\u0010\u0019\u001a\u00020\u000e2\u0006\u0010\u001a\u001a\u00020\u0018H\u0002J\u0010\u0010\u001b\u001a\u00020\u000e2\u0006\u0010\u0017\u001a\u00020\u0018H\u0002J\u0010\u0010\u001c\u001a\u00020\u000e2\u0006\u0010\u001a\u001a\u00020\u0018H\u0002J\b\u0010\u001d\u001a\u00020\u000eH\u0002J\b\u0010\u001e\u001a\u00020\u000eH\u0002J\b\u0010\u001f\u001a\u00020\u000eH\u0002J\u0012\u0010 \u001a\u00020\u000e2\b\u0010!\u001a\u0004\u0018\u00010\"H\u0014J\u0006\u0010#\u001a\u00020\u000eJ\b\u0010$\u001a\u00020\u000eH\u0002R\u000e\u0010\u0003\u001a\u00020\u0004X\u0082.\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0005\u001a\u00020\u0006X\u0082.\u00a2\u0006\u0002\n\u0000R\u001e\u0010\u0007\u001a\u00020\b8\u0006@\u0006X\u0087.\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\t\u0010\n\"\u0004\b\u000b\u0010\f\u00a8\u0006&"}, d2 = {"Lorg/csystem/android/app/generator/random/MainActivity;", "Landroidx/appcompat/app/AppCompatActivity;", "()V", "mBinding", "Lorg/csystem/android/app/generator/random/databinding/ActivityMainBinding;", "mHandler", "Landroid/os/Handler;", "threadPool", "Ljava/util/concurrent/ExecutorService;", "getThreadPool", "()Ljava/util/concurrent/ExecutorService;", "setThreadPool", "(Ljava/util/concurrent/ExecutorService;)V", "clearButtonClicked", "", "configurationButtonClicked", "configurationCallback", "configurationConnectionCallback", "socket", "Ljava/net/Socket;", "generateButtonClicked", "generatorConnectionCallback", "handleAnyException", "message", "", "handleGetText", "text", "handleIOException", "handleInvalidValues", "initBinding", "initViewModels", "initialize", "onCreate", "savedInstanceState", "Landroid/os/Bundle;", "saveButtonClicked", "textGeneratorCallback", "GetTextHandler", "app_debug"})
public final class MainActivity extends androidx.appcompat.app.AppCompatActivity {
    private org.csystem.android.app.generator.random.databinding.ActivityMainBinding mBinding;
    private android.os.Handler mHandler;
    @javax.inject.Inject
    public java.util.concurrent.ExecutorService threadPool;
    
    public MainActivity() {
        super();
    }
    
    @org.jetbrains.annotations.NotNull
    public final java.util.concurrent.ExecutorService getThreadPool() {
        return null;
    }
    
    public final void setThreadPool(@org.jetbrains.annotations.NotNull
    java.util.concurrent.ExecutorService p0) {
    }
    
    private final void handleGetText(java.lang.String text) {
    }
    
    private final void handleInvalidValues(java.lang.String text) {
    }
    
    private final void handleIOException(java.lang.String message) {
    }
    
    private final void handleAnyException(java.lang.String message) {
    }
    
    private final void generatorConnectionCallback(java.net.Socket socket) {
    }
    
    private final void configurationConnectionCallback(java.net.Socket socket) {
    }
    
    private final void textGeneratorCallback() {
    }
    
    private final void configurationCallback() {
    }
    
    private final void initViewModels() {
    }
    
    private final void initBinding() {
    }
    
    private final void initialize() {
    }
    
    @java.lang.Override
    protected void onCreate(@org.jetbrains.annotations.Nullable
    android.os.Bundle savedInstanceState) {
    }
    
    public final void generateButtonClicked() {
    }
    
    public final void configurationButtonClicked() {
    }
    
    public final void saveButtonClicked() {
    }
    
    public final void clearButtonClicked() {
    }
    
    @kotlin.Metadata(mv = {1, 9, 0}, k = 1, xi = 48, d1 = {"\u0000&\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0010\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\b\u0002\u0018\u00002\u00020\u0001B\r\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u00a2\u0006\u0002\u0010\u0004J\u0010\u0010\b\u001a\u00020\t2\u0006\u0010\n\u001a\u00020\u000bH\u0016R\u001c\u0010\u0005\u001a\u0010\u0012\f\u0012\n \u0007*\u0004\u0018\u00010\u00030\u00030\u0006X\u0082\u0004\u00a2\u0006\u0002\n\u0000\u00a8\u0006\f"}, d2 = {"Lorg/csystem/android/app/generator/random/MainActivity$GetTextHandler;", "Landroid/os/Handler;", "activity", "Lorg/csystem/android/app/generator/random/MainActivity;", "(Lorg/csystem/android/app/generator/random/MainActivity;)V", "mWeakReference", "Ljava/lang/ref/WeakReference;", "kotlin.jvm.PlatformType", "handleMessage", "", "msg", "Landroid/os/Message;", "app_debug"})
    static final class GetTextHandler extends android.os.Handler {
        @org.jetbrains.annotations.NotNull
        private final java.lang.ref.WeakReference<org.csystem.android.app.generator.random.MainActivity> mWeakReference = null;
        
        public GetTextHandler(@org.jetbrains.annotations.NotNull
        org.csystem.android.app.generator.random.MainActivity activity) {
            super();
        }
        
        @java.lang.Override
        public void handleMessage(@org.jetbrains.annotations.NotNull
        android.os.Message msg) {
        }
    }
}